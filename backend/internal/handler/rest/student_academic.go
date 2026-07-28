package rest

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"dbms-project/internal/db"
	"dbms-project/internal/middleware"
)

type StudentAcademicPayload struct {
	ExamLevel string `json:"exam_level"`
	Year      int32  `json:"year"`
	RollNo    string `json:"roll_no"`
	RegNo     string `json:"reg_no"`
	Gpa       string `json:"gpa"`
	Board     string `json:"board"`
	EduGroup  string `json:"edu_group"` // 'Science', 'Humanities', 'Business', etc.
}

type SubjectMarkPayload struct {
	SubjectName string `json:"subject_name"`
	Marks       string `json:"marks"`
	Grade       string `json:"grade"`
}

type StudentSubjectMarksPayload struct {
	ExamLevel string               `json:"exam_level"`
	Subjects  []SubjectMarkPayload `json:"subjects"`
}

func (h *Handler) HandleAddStudentAcademic(w http.ResponseWriter, r *http.Request) {
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req StudentAcademicPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	examLevel := strings.TrimSpace(req.ExamLevel)
	rollNo := strings.TrimSpace(req.RollNo)
	regNo := strings.TrimSpace(req.RegNo)
	board := strings.TrimSpace(req.Board)
	eduGroup := strings.TrimSpace(req.EduGroup)

	if examLevel == "" || rollNo == "" || regNo == "" || board == "" || req.Year <= 0 {
		http.Error(w, "exam_level, year, roll_no, reg_no, and board are required", http.StatusBadRequest)
		return
	}

	gpaVal, err := strconv.ParseFloat(req.Gpa, 64)
	if err != nil || gpaVal < 0.0 || gpaVal > 5.0 {
		http.Error(w, "GPA must be a valid number between 0.00 and 5.00", http.StatusBadRequest)
		return
	}

	if eduGroup == "" {
		eduGroup = "Science" // Default
	}

	err = h.Queries.AddStudentAcademic(r.Context(), db.AddStudentAcademicParams{
		StudentID: studentID,
		ExamLevel: examLevel,
		Year:      req.Year,
		RollNo:    rollNo,
		RegNo:     regNo,
		Gpa:       req.Gpa,
		Board:     board,
		EduGroup:  eduGroup,
	})
	if err != nil {
		http.Error(w, "Failed to save academic records: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Academic records saved successfully",
	})
}

func (h *Handler) HandleAddStudentSubjectMarks(w http.ResponseWriter, r *http.Request) {
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req StudentSubjectMarksPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	examLevel := strings.TrimSpace(req.ExamLevel)
	if examLevel == "" {
		http.Error(w, "exam_level is required", http.StatusBadRequest)
		return
	}

	// Start a transaction to insert/update subject marks atomically
	tx, err := h.DB.BeginTx(r.Context(), nil)
	if err != nil {
		http.Error(w, "Failed to start database transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	qtx := h.Queries.WithTx(tx)

	for _, sub := range req.Subjects {
		subName := strings.TrimSpace(sub.SubjectName)
		marksStr := strings.TrimSpace(sub.Marks)
		gradeStr := strings.TrimSpace(sub.Grade)

		if subName == "" || marksStr == "" {
			http.Error(w, "subject_name and marks are required for all elements", http.StatusBadRequest)
			return
		}

		marksVal, err := strconv.ParseFloat(marksStr, 64)
		if err != nil || marksVal < 0 || marksVal > 100 {
			http.Error(w, "Subject marks must be a valid number between 0.00 and 100.00", http.StatusBadRequest)
			return
		}

		var gradeVal sql.NullString
		if gradeStr != "" {
			gradeVal = sql.NullString{String: gradeStr, Valid: true}
		}

		err = qtx.UpsertStudentSubjectMark(r.Context(), db.UpsertStudentSubjectMarkParams{
			StudentID:   studentID,
			ExamLevel:   examLevel,
			SubjectName: subName,
			Marks:       marksStr,
			Grade:       gradeVal,
		})
		if err != nil {
			http.Error(w, "Failed to save subject mark for "+subName+": "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to save subject marks transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Subject marks saved successfully",
	})
}

func (h *Handler) HandleGetStudentNotifications(w http.ResponseWriter, r *http.Request) {
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	notifs, err := h.Queries.GetStudentNotifications(r.Context(), studentID)
	if err != nil {
		http.Error(w, "Failed to fetch notifications: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if notifs == nil {
		notifs = []db.Notification{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notifs)
}
