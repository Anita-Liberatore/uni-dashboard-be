package domain

import "context"

// StudentRepository is the outbound port (driven side) for student profile data.
type StudentRepository interface {
	GetProfile(ctx context.Context) (*Student, error)
	GetAcademicRecord(ctx context.Context) (*AcademicRecord, error)
}

// ExamRepository is the outbound port for exam data.
type ExamRepository interface {
	GetPassed(ctx context.Context) ([]Exam, error)
	GetUpcoming(ctx context.Context) ([]UpcomingExam, error)
}

// StudyPlanRepository is the outbound port for the student's study plan.
type StudyPlanRepository interface {
	GetStudyPlan(ctx context.Context) ([]YearPlan, error)
}

// DocumentRepository is the outbound port for uploaded documents.
type DocumentRepository interface {
	GetDocuments(ctx context.Context) ([]Document, error)
}

// CalendarRepository is the outbound port for calendar events.
type CalendarRepository interface {
	GetEvents(ctx context.Context) ([]CalendarEvent, error)
}
