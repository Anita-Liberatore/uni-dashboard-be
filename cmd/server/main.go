package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/liberatoreanita/uni-dashboard-be/internal/adapters/memory"
	httpadapter "github.com/liberatoreanita/uni-dashboard-be/internal/adapters/http"
	"github.com/liberatoreanita/uni-dashboard-be/internal/application"
)

const addr = ":8080"

func main() {
	// ── Driven adapters (right side) ─────────────────────────────────────────
	// Swap any of these with a real DB implementation without touching anything else.
	studentRepo  := memory.NewStudentRepository()
	examRepo     := memory.NewExamRepository()
	studyPlanRepo := memory.NewStudyPlanRepository()
	documentRepo := memory.NewDocumentRepository()
	calendarRepo := memory.NewCalendarRepository()

	// ── Application layer ────────────────────────────────────────────────────
	svc := application.NewStudentService(
		studentRepo,
		examRepo,
		studyPlanRepo,
		documentRepo,
		calendarRepo,
	)

	// ── Driving adapter (left side) ──────────────────────────────────────────
	handler := httpadapter.NewHandler(svc)
	router  := httpadapter.NewRouter(handler)

	// ── Server ───────────────────────────────────────────────────────────────
	fmt.Printf("uni-dashboard-be listening on %s\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
