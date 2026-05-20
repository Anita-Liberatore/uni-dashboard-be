// Package application contains the use cases for the student dashboard.
// It depends only on domain ports — no HTTP, no database, no framework.
package application

import (
	"context"

	"github.com/liberatoreanita/uni-dashboard-be/internal/domain"
)

// StudentService orchestrates all use cases for the student dashboard.
// It is injected with outbound ports (repositories) at startup.
type StudentService struct {
	students  domain.StudentRepository
	exams     domain.ExamRepository
	studyPlan domain.StudyPlanRepository
	documents domain.DocumentRepository
	calendar  domain.CalendarRepository
}

// NewStudentService constructs the service with all required ports.
func NewStudentService(
	students domain.StudentRepository,
	exams domain.ExamRepository,
	studyPlan domain.StudyPlanRepository,
	documents domain.DocumentRepository,
	calendar domain.CalendarRepository,
) *StudentService {
	return &StudentService{
		students:  students,
		exams:     exams,
		studyPlan: studyPlan,
		documents: documents,
		calendar:  calendar,
	}
}

func (s *StudentService) GetProfile(ctx context.Context) (*domain.Student, error) {
	return s.students.GetProfile(ctx)
}

func (s *StudentService) GetAcademicRecord(ctx context.Context) (*domain.AcademicRecord, error) {
	return s.students.GetAcademicRecord(ctx)
}

func (s *StudentService) GetExamsPassed(ctx context.Context) ([]domain.Exam, error) {
	return s.exams.GetPassed(ctx)
}

func (s *StudentService) GetExamsUpcoming(ctx context.Context) ([]domain.UpcomingExam, error) {
	return s.exams.GetUpcoming(ctx)
}

func (s *StudentService) GetStudyPlan(ctx context.Context) ([]domain.YearPlan, error) {
	return s.studyPlan.GetStudyPlan(ctx)
}

func (s *StudentService) GetDocuments(ctx context.Context) ([]domain.Document, error) {
	return s.documents.GetDocuments(ctx)
}

func (s *StudentService) GetCalendarEvents(ctx context.Context) ([]domain.CalendarEvent, error) {
	return s.calendar.GetEvents(ctx)
}
