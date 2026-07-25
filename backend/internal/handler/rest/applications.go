package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"dbms-project/internal/middleware"
)

type EligibleProgram struct {
	ProgramID      int32  `json:"program_id"`
	ProgramName    string `json:"program_name"`
	UniversityName string `json:"university_name"`
}

func (h *Handler) HandleEligiblePrograms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Get logged-in student's ID from context
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ctx := r.Context()

	// 2. Fetch Student's Academics using your GetStudentAcademics query
	academics, err := h.Queries.GetStudentAcademics(ctx, studentID)
	if err != nil {
		http.Error(w, "Failed to fetch student academics", http.StatusInternalServerError)
		return
	}

	// Map exam_level (e.g. "SSC", "HSC") to GPA float
	studentGPA := make(map[string]float64)
	for _, acad := range academics {
		// acad.ExamLevel is string, acad.Gpa is string (NUMERIC)
		gpaVal, err := strconv.ParseFloat(acad.Gpa, 64)
		if err == nil {
			studentGPA[acad.ExamLevel] = gpaVal
		}
	}

	// 3. Fetch all programs and their eligibility rules
	programRules, err := h.Queries.GetAllProgramsWithRules(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch programs", http.StatusInternalServerError)
		return
	}

	// 4. Evaluate eligibility logic
	eligibleMap := make(map[int32]EligibleProgram)
	disqualifiedMap := make(map[int32]bool)

	for _, pr := range programRules {
		if disqualifiedMap[pr.ProgramID] {
			continue
		}

		// Check rules if rule_type and rule_value are present
		if pr.RuleType.Valid && pr.RuleValue.Valid {
			requiredGPA, _ := strconv.ParseFloat(pr.RuleValue.String, 64)

			if pr.RuleType.String == "MIN_SSC_GPA" && studentGPA["SSC"] < requiredGPA {
				disqualifiedMap[pr.ProgramID] = true
				delete(eligibleMap, pr.ProgramID)
				continue
			}
			if pr.RuleType.String == "MIN_HSC_GPA" && studentGPA["HSC"] < requiredGPA {
				disqualifiedMap[pr.ProgramID] = true
				delete(eligibleMap, pr.ProgramID)
				continue
			}
		}

		// Add program to eligible map if not disqualified
		eligibleMap[pr.ProgramID] = EligibleProgram{
			ProgramID:      pr.ProgramID,
			ProgramName:    pr.PName,
			UniversityName: pr.UniversityName,
		}
	}

	// 5. Format results into JSON array
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
