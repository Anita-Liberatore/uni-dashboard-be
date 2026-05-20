package httpadapter

import "net/http"

// NewRouter wires all routes onto a standard ServeMux (Go 1.22+).
// Method-qualified patterns ("GET /path") prevent accidental wrong-method hits.
func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/me",                 h.GetProfile)
	mux.HandleFunc("GET /api/v1/me/academic",        h.GetAcademic)
	mux.HandleFunc("GET /api/v1/me/exams",           h.GetExamsPassed)
	mux.HandleFunc("GET /api/v1/me/exams/upcoming",  h.GetExamsUpcoming)
	mux.HandleFunc("GET /api/v1/me/study-plan",      h.GetStudyPlan)
	mux.HandleFunc("GET /api/v1/me/documents",       h.GetDocuments)
	mux.HandleFunc("GET /api/v1/me/calendar",        h.GetCalendar)

	return withCORS(mux)
}

// withCORS adds permissive CORS headers for local development.
// In production, replace the wildcard origin with the actual FE domain.
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
