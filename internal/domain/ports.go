package domain

import "context"

// StudentRepository is the outbound port (driven side) for student profile data.
type StudentRepository interface {
	GetProfile(ctx context.Context, studentID string) (*Student, error)
	GetAcademicRecord(ctx context.Context, studentID string) (*AcademicRecord, error)
}

// ExamRepository is the outbound port for exam data.
type ExamRepository interface {
	GetPassed(ctx context.Context, studentID string) ([]Exam, error)
	GetUpcoming(ctx context.Context, studentID string) ([]UpcomingExam, error)
}

// StudyPlanRepository is the outbound port for the student's study plan.
type StudyPlanRepository interface {
	GetStudyPlan(ctx context.Context, studentID string) ([]YearPlan, error)
}

// DocumentRepository is the outbound port for uploaded documents.
type DocumentRepository interface {
	GetDocuments(ctx context.Context, studentID string) ([]Document, error)
}

// CalendarRepository is the outbound port for calendar events.
type CalendarRepository interface {
	GetEvents(ctx context.Context, studentID string) ([]CalendarEvent, error)
}
