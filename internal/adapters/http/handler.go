// Package httpadapter is the driving (inbound) adapter.
// It translates HTTP requests into application use-case calls and writes JSON responses.
package httpadapter

import (
	"encoding/json"
	"net/http"

	"github.com/liberatoreanita/uni-dashboard-be/internal/application"
	"github.com/liberatoreanita/uni-dashboard-be/internal/domain"
)

// Handler holds a reference to the application service.
// It knows nothing about persistence — only about HTTP and JSON.
type Handler struct {
	svc *application.StudentService
}

// NewHandler constructs the HTTP handler with the required service.
func NewHandler(svc *application.StudentService) *Handler {
	return &Handler{svc: svc}
}

// ErrorResponse is the standard error payload returned by all endpoints.
type ErrorResponse struct {
	Error string `json:"error" example:"something went wrong"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Helpers
// ─────────────────────────────────────────────────────────────────────────────

func (h *Handler) writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func (h *Handler) writeError(w http.ResponseWriter, status int, err error) {
	h.writeJSON(w, status, ErrorResponse{Error: err.Error()})
}

// ─────────────────────────────────────────────────────────────────────────────
// Endpoints
// ─────────────────────────────────────────────────────────────────────────────

// GetStudent handles GET /api/v1/students/{studentId}
//
//	@Summary		Get student profile
//	@Description	Returns the complete profile for the given student
//	@Tags			students
//	@Produce		json
//	@Param			studentId	path		string			true	"Student ID (e.g. S1234567)"
//	@Success		200			{object}	domain.Student
//	@Failure		500			{object}	ErrorResponse
//	@Router			/students/{studentId} [get]
func (h *Handler) GetStudent(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("studentId")
	data, err := h.svc.GetProfile(r.Context(), id)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}

// GetAcademic handles GET /api/v1/students/{studentId}/academic
//
//	@Summary		Get academic record
//	@Description	Returns GPA, credits, courses progress and degree details
//	@Tags			students
//	@Produce		json
//	@Param			studentId	path		string					true	"Student ID (e.g. S1234567)"
//	@Success		200			{object}	domain.AcademicRecord
//	@Failure		500			{object}	ErrorResponse
//	@Router			/students/{studentId}/academic [get]
func (h *Handler) GetAcademic(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("studentId")
	data, err := h.svc.GetAcademicRecord(r.Context(), id)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}

// GetExamsPassed handles GET /api/v1/students/{studentId}/exams
//
//	@Summary		List passed exams
//	@Description	Returns the full list of exams passed by the student, ordered by date desc
//	@Tags			exams
//	@Produce		json
//	@Param			studentId	path		string			true	"Student ID (e.g. S1234567)"
//	@Success		200			{array}		domain.Exam
//	@Failure		500			{object}	ErrorResponse
//	@Router			/students/{studentId}/exams [get]
func (h *Handler) GetExamsPassed(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("studentId")
	data, err := h.svc.GetExamsPassed(r.Context(), id)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}

// GetExamsUpcoming handles GET /api/v1/students/{studentId}/exams/upcoming
//
//	@Summary		List upcoming exams
//	@Description	Returns exams the student is yet to sit, ordered by date asc
//	@Tags			exams
//	@Produce		json
//	@Param			studentId	path		string					true	"Student ID (e.g. S1234567)"
//	@Success		200			{array}		domain.UpcomingExam
//	@Failure		500			{object}	ErrorResponse
//	@Router			/students/{studentId}/exams/upcoming [get]
func (h *Handler) GetExamsUpcoming(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("studentId")
	data, err := h.svc.GetExamsUpcoming(r.Context(), id)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}

// GetStudyPlan handles GET /api/v1/students/{studentId}/study-plan
//
//	@Summary		Get study plan
//	@Description	Returns the full study plan grouped by academic year
//	@Tags			students
//	@Produce		json
//	@Param			studentId	path		string			true	"Student ID (e.g. S1234567)"
//	@Success		200			{array}		domain.YearPlan
//	@Failure		500			{object}	ErrorResponse
//	@Router			/students/{studentId}/study-plan [get]
func (h *Handler) GetStudyPlan(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("studentId")
	data, err := h.svc.GetStudyPlan(r.Context(), id)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}

// GetDocuments handles GET /api/v1/students/{studentId}/documents
//
//	@Summary		List documents
//	@Description	Returns all documents uploaded by the student
//	@Tags			students
//	@Produce		json
//	@Param			studentId	path		string				true	"Student ID (e.g. S1234567)"
//	@Success		200			{array}		domain.Document
//	@Failure		500			{object}	ErrorResponse
//	@Router			/students/{studentId}/documents [get]
func (h *Handler) GetDocuments(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("studentId")
	data, err := h.svc.GetDocuments(r.Context(), id)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}

// GetCalendar handles GET /api/v1/students/{studentId}/calendar
//
//	@Summary		Get calendar events
//	@Description	Returns scheduled events (exams, deadlines, lectures) ordered by date asc
//	@Tags			students
//	@Produce		json
//	@Param			studentId	path		string					true	"Student ID (e.g. S1234567)"
//	@Success		200			{array}		domain.CalendarEvent
//	@Failure		500			{object}	ErrorResponse
//	@Router			/students/{studentId}/calendar [get]
func (h *Handler) GetCalendar(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("studentId")
	data, err := h.svc.GetCalendarEvents(r.Context(), id)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}

// ensure domain types are visible to the swag parser for annotation resolution
var _ = domain.Student{}
