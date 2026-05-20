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

func (s *StudentService) GetProfile(ctx context.Context, studentID string) (*domain.Student, error) {
	return s.students.GetProfile(ctx, studentID)
}

func (s *StudentService) GetAcademicRecord(ctx context.Context, studentID string) (*domain.AcademicRecord, error) {
	return s.students.GetAcademicRecord(ctx, studentID)
}

func (s *StudentService) GetExamsPassed(ctx context.Context, studentID string) ([]domain.Exam, error) {
	return s.exams.GetPassed(ctx, studentID)
}

func (s *StudentService) GetExamsUpcoming(ctx context.Context, studentID string) ([]domain.UpcomingExam, error) {
	return s.exams.GetUpcoming(ctx, studentID)
}

func (s *StudentService) GetStudyPlan(ctx context.Context, studentID string) ([]domain.YearPlan, error) {
	return s.studyPlan.GetStudyPlan(ctx, studentID)
}

func (s *StudentService) GetDocuments(ctx context.Context, studentID string) ([]domain.Document, error) {
	return s.documents.GetDocuments(ctx, studentID)
}

func (s *StudentService) GetCalendarEvents(ctx context.Context, studentID string) ([]domain.CalendarEvent, error) {
	return s.calendar.GetEvents(ctx, studentID)
}
