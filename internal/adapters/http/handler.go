// Package httpadapter is the driving (inbound) adapter.
// It translates HTTP requests into application use-case calls and writes JSON responses.
package httpadapter

import (
	"encoding/json"
	"net/http"

	"github.com/liberatoreanita/uni-dashboard-be/internal/application"
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

// ─────────────────────────────────────────────────────────────────────────────
// Helpers
// ─────────────────────────────────────────────────────────────────────────────

func (h *Handler) writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func (h *Handler) writeError(w http.ResponseWriter, status int, err error) {
	h.writeJSON(w, status, map[string]string{"error": err.Error()})
}

// ─────────────────────────────────────────────────────────────────────────────
// Endpoints — each maps 1-to-1 to a use case
// ─────────────────────────────────────────────────────────────────────────────

// GetProfile handles GET /api/v1/me
func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	data, err := h.svc.GetProfile(r.Context())
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}

// GetAcademic handles GET /api/v1/me/academic
func (h *Handler) GetAcademic(w http.ResponseWriter, r *http.Request) {
	data, err := h.svc.GetAcademicRecord(r.Context())
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}

// GetExamsPassed handles GET /api/v1/me/exams
func (h *Handler) GetExamsPassed(w http.ResponseWriter, r *http.Request) {
	data, err := h.svc.GetExamsPassed(r.Context())
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}

// GetExamsUpcoming handles GET /api/v1/me/exams/upcoming
func (h *Handler) GetExamsUpcoming(w http.ResponseWriter, r *http.Request) {
	data, err := h.svc.GetExamsUpcoming(r.Context())
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}

// GetStudyPlan handles GET /api/v1/me/study-plan
func (h *Handler) GetStudyPlan(w http.ResponseWriter, r *http.Request) {
	data, err := h.svc.GetStudyPlan(r.Context())
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}

// GetDocuments handles GET /api/v1/me/documents
func (h *Handler) GetDocuments(w http.ResponseWriter, r *http.Request) {
	data, err := h.svc.GetDocuments(r.Context())
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}

// GetCalendar handles GET /api/v1/me/calendar
func (h *Handler) GetCalendar(w http.ResponseWriter, r *http.Request) {
	data, err := h.svc.GetCalendarEvents(r.Context())
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err)
		return
	}
	h.writeJSON(w, http.StatusOK, data)
}
