package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"dbms-project/internal/db"
	"dbms-project/internal/middleware"
)

type EligibleProgram struct {
	ProgramID      int32  `json:"program_id"`
	ProgramName    string `json:"program_name"`
	UniversityName string `json:"university_name"`
}

func (h *Handler) HandleEligiblePrograms(w http.ResponseWriter, r *http.Request) {
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ctx := r.Context()

	// 1. Fetch Student Academics
	academics, err := h.Queries.GetStudentAcademics(ctx, studentID)
	if err != nil {
		http.Error(w, "Failed to fetch student academics", http.StatusInternalServerError)
		return
	}

	studentGPA := make(map[string]float64)
	studentGroup := make(map[string]string)

	for _, acad := range academics {
		if gpaVal, err := strconv.ParseFloat(acad.Gpa, 64); err == nil {
			studentGPA[acad.ExamLevel] = gpaVal
		}
		studentGroup[acad.ExamLevel] = acad.EduGroup // 'Science', 'Humanities', etc.
	}

	// 2. Fetch HSC Subject Marks (Crucial for Engineering & Unit A)
	hscMarksList, err := h.Queries.GetStudentSubjectMarks(ctx, db.GetStudentSubjectMarksParams{
		StudentID: studentID,
		ExamLevel: "HSC",
	})
	if err != nil {
		http.Error(w, "Failed to fetch subject marks", http.StatusInternalServerError)
		return
	}

	hscMarks := make(map[string]float64)
	for _, m := range hscMarksList {
		if val, err := strconv.ParseFloat(m.Marks, 64); err == nil {
			hscMarks[m.SubjectName] = val
		}
	}

	// 3. Fetch Programs & Rules
	programRules, err := h.Queries.GetAllProgramsWithRules(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch programs", http.StatusInternalServerError)
		return
	}

	// 4. Evaluate Rules
	eligibleMap := make(map[int32]EligibleProgram)
	disqualifiedMap := make(map[int32]bool)

	for _, pr := range programRules {
		if disqualifiedMap[pr.ProgramID] {
			continue
		}

		if pr.RuleType.Valid && pr.RuleValue.Valid {
			ruleType := pr.RuleType.String
			ruleVal := pr.RuleValue.String

			// Check Required Education Group (e.g. Science only)
			if ruleType == "REQ_HSC_GROUP" && studentGroup["HSC"] != ruleVal {
				disqualifiedMap[pr.ProgramID] = true
				delete(eligibleMap, pr.ProgramID)
				continue
			}

			// Check Minimum GPAs
			if requiredGPA, err := strconv.ParseFloat(ruleVal, 64); err == nil {
				if ruleType == "MIN_SSC_GPA" && studentGPA["SSC"] < requiredGPA {
					disqualifiedMap[pr.ProgramID] = true
					delete(eligibleMap, pr.ProgramID)
					continue
				}
				if ruleType == "MIN_HSC_GPA" && studentGPA["HSC"] < requiredGPA {
					disqualifiedMap[pr.ProgramID] = true
					delete(eligibleMap, pr.ProgramID)
					continue
				}
			}

			// Check Subject Specific Marks (e.g. MIN_HSC_PHYSICS = 80)
			if ruleType == "MIN_HSC_PHYSICS" || ruleType == "MIN_HSC_MATH" || ruleType == "MIN_HSC_CHEMISTRY" {
				reqScore, _ := strconv.ParseFloat(ruleVal, 64)
				subj := "Physics"
				switch ruleType {
				case "MIN_HSC_MATH":
					subj = "Mathematics"
				case "MIN_HSC_CHEMISTRY":
					subj = "Chemistry"
				}

				if hscMarks[subj] < reqScore {
					disqualifiedMap[pr.ProgramID] = true
					delete(eligibleMap, pr.ProgramID)
					continue
				}
			}
		}

		eligibleMap[pr.ProgramID] = EligibleProgram{
			ProgramID:      pr.ProgramID,
			ProgramName:    pr.PName,
			UniversityName: pr.UniversityName,
		}
	}

	var results []EligibleProgram
	for _, p := range eligibleMap {
		results = append(results, p)
	}
	if results == nil {
		results = []EligibleProgram{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

type ApplyRequest struct {
	ProgramID int32 `json:"program_id"`
}

func (h *Handler) ApplyToProgram(w http.ResponseWriter, r *http.Request) {
	// 1. Get Student ID from JWT Context using your middleware key
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req ApplyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	// 2. Fetch what the program requires (Changed h.DB to h.Queries)
	requiredFields, err := h.Queries.GetProgramRequiredFields(ctx, req.ProgramID)
	if err != nil {
		http.Error(w, "Error fetching program requirements", http.StatusInternalServerError)
		return
	}

	// 3. Fetch what the student has already provided (Changed h.DB to h.Queries)
	studentFields, err := h.Queries.GetStudentProfileFields(ctx, studentID)
	if err != nil {
		http.Error(w, "Error fetching student profile", http.StatusInternalServerError)
		return
	}

	// 4. Calculate Missing Fields
	// Convert student fields into a map for fast lookup
	providedMap := make(map[string]bool)
	for _, field := range studentFields {
		providedMap[string(field)] = true
	}

	var missingFields []string
	for _, reqField := range requiredFields {
		if !providedMap[string(reqField)] {
			missingFields = append(missingFields, string(reqField))
		}
	}

	// 5. If fields are missing, reject the application and tell them what to fill out
	if len(missingFields) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":         "incomplete_profile",
			"message":        "You must provide additional information to apply to this program.",
			"missing_fields": missingFields,
		})
		return
	}

	// 6. If no fields are missing, insert the application! (Changed h.DB to h.Queries)
	res, err := h.Queries.CreateApplication(ctx, db.CreateApplicationParams{
		StudentID: studentID,
		ProgramID: req.ProgramID,
	})
	if err != nil {
		http.Error(w, "Failed to submit application", http.StatusInternalServerError)
		return
	}

	// Extract the newly generated auto-increment app_id
	appID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve application ID", http.StatusInternalServerError)
		return
	}

	// Success!
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  "success",
		"message": "Application submitted successfully! Please proceed to payment.",
		"app_id":  appID, // <-- Returns the real ID!
	})
}

func (h *Handler) HandleGetStudentApplications(w http.ResponseWriter, r *http.Request) {
	// 1. Authenticate Student
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ctx := r.Context()

	// 2. Fetch applications from database
	applications, err := h.Queries.GetStudentApplications(ctx, studentID)
	if err != nil {
		http.Error(w, "Failed to fetch applications", http.StatusInternalServerError)
		return
	}

	// Ensure we return an empty JSON array `[]` instead of `null` if no records exist
	if applications == nil {
		applications = []db.GetStudentApplicationsRow{}
	}

	// 3. Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(applications)
}
