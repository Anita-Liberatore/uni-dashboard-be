// Package memory provides in-memory implementations of all domain repository ports.
// Replace with a real database adapter (e.g. postgres/) without touching any other layer.
package memory

import (
	"context"

	"github.com/liberatoreanita/uni-dashboard-be/internal/domain"
)

// ─────────────────────────────────────────────────────────────────────────────
// Helpers
// ─────────────────────────────────────────────────────────────────────────────

func ptr(i int) *int { return &i }

// ─────────────────────────────────────────────────────────────────────────────
// StudentRepository
// ─────────────────────────────────────────────────────────────────────────────

type StudentRepository struct{}

func NewStudentRepository() *StudentRepository { return &StudentRepository{} }

func (r *StudentRepository) GetProfile(_ context.Context) (*domain.Student, error) {
	s := domain.Student{
		Name:           "Anita",
		Surname:        "Liberatore",
		Area:           "Ingegneria Informatica",
		StudentID:      "S1234567",
		Email:          "anita.liberatore@studenti.unito.it",
		PEC:            "anita.liberatore@pec.unito.it",
		Address:        "Via Roma 123, 10100 Torino (TO)",
		Phone:          "+39 320 123 4567",
		BirthDate:      "1998-04-22",
		BirthPlace:     "Napoli (NA)",
		TaxCode:        "LBRANT98D62F839X",
		Status:         "ACTIVE",
		Year:           3,
		Semester:       1,
		EnrolledSince:  "1 set 2021",
		GraduationDate: "giu 2025",
		Advisor:        "Prof.ssa Laura Bianchi",
	}
	return &s, nil
}

func (r *StudentRepository) GetAcademicRecord(_ context.Context) (*domain.AcademicRecord, error) {
	rec := domain.AcademicRecord{
		Program:   "Ingegneria Informatica",
		Degree:    "Laurea Triennale (L-8)",
		Faculty:   "Facoltà di Scienze e Tecnologia",
		Credits:   domain.ProgressPair{Current: 53, Total: 180},
		Courses:   domain.ProgressPair{Current: 6, Total: 22},
		Electives: domain.ProgressPair{Current: 0, Total: 4},
		GPA:       29.8,
	}
	return &rec, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// ExamRepository
// ─────────────────────────────────────────────────────────────────────────────

type ExamRepository struct{}

func NewExamRepository() *ExamRepository { return &ExamRepository{} }

func (r *ExamRepository) GetPassed(_ context.Context) ([]domain.Exam, error) {
	return []domain.Exam{
		{Course: "Algoritmi e Strutture Dati", Area: "Inf.", Date: "Jan 18, 2024", Grade: 30, Lode: true, Credits: 9, Year: 3},
		{Course: "Basi di Dati", Area: "Inf.", Date: "Jun 20, 2023", Grade: 28, Lode: false, Credits: 9, Year: 2},
		{Course: "Reti di Calcolatori", Area: "Reti", Date: "Jan 12, 2023", Grade: 27, Lode: false, Credits: 6, Year: 2},
		{Course: "Programmazione II", Area: "Inf.", Date: "Feb 15, 2023", Grade: 30, Lode: true, Credits: 9, Year: 2},
		{Course: "Analisi Matematica II", Area: "Mat.", Date: "Jul 10, 2022", Grade: 26, Lode: false, Credits: 12, Year: 1},
		{Course: "Programmazione I", Area: "Inf.", Date: "Jan 20, 2022", Grade: 30, Lode: true, Credits: 9, Year: 1},
	}, nil
}

func (r *ExamRepository) GetUpcoming(_ context.Context) ([]domain.UpcomingExam, error) {
	return []domain.UpcomingExam{
		{Course: "Sistemi Operativi", Date: "May 28, 2024", Credits: 9, Urgent: true},
		{Course: "Ingegneria del Software", Date: "Jun 15, 2024", Credits: 9, Urgent: false},
		{Course: "Intelligenza Artificiale", Date: "Jul 10, 2024", Credits: 6, Urgent: false},
	}, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// StudyPlanRepository
// ─────────────────────────────────────────────────────────────────────────────

type StudyPlanRepository struct{}

func NewStudyPlanRepository() *StudyPlanRepository { return &StudyPlanRepository{} }

func (r *StudyPlanRepository) GetStudyPlan(_ context.Context) ([]domain.YearPlan, error) {
	return []domain.YearPlan{
		{
			Year: 1,
			Courses: []domain.Course{
				{Name: "Analisi Matematica I", Credits: 12, Semester: 1, Passed: true, Grade: ptr(26)},
				{Name: "Programmazione I", Credits: 9, Semester: 1, Passed: true, Grade: ptr(30), Lode: true},
				{Name: "Algebra Lineare", Credits: 9, Semester: 1, Passed: true, Grade: ptr(27)},
				{Name: "Analisi Matematica II", Credits: 12, Semester: 2, Passed: true, Grade: ptr(26)},
				{Name: "Fisica I", Credits: 9, Semester: 2, Passed: false},
				{Name: "Architettura degli Elaboratori", Credits: 6, Semester: 2, Passed: false},
			},
		},
		{
			Year: 2,
			Courses: []domain.Course{
				{Name: "Programmazione II", Credits: 9, Semester: 1, Passed: true, Grade: ptr(30), Lode: true},
				{Name: "Basi di Dati", Credits: 9, Semester: 1, Passed: true, Grade: ptr(28)},
				{Name: "Reti di Calcolatori", Credits: 6, Semester: 1, Passed: true, Grade: ptr(27)},
				{Name: "Sistemi Operativi", Credits: 9, Semester: 2, Passed: false},
				{Name: "Ingegneria del Software", Credits: 9, Semester: 2, Passed: false},
				{Name: "Statistica", Credits: 6, Semester: 2, Passed: false},
			},
		},
		{
			Year: 3,
			Courses: []domain.Course{
				{Name: "Algoritmi e Strutture Dati", Credits: 9, Semester: 1, Passed: true, Grade: ptr(30), Lode: true},
				{Name: "Intelligenza Artificiale", Credits: 6, Semester: 1, Passed: false},
				{Name: "Sicurezza Informatica", Credits: 6, Semester: 1, Passed: false},
				{Name: "Tesi di Laurea I", Credits: 6, Semester: 2, Passed: false},
				{Name: "Corso a Scelta I", Credits: 6, Semester: 2, Passed: false},
				{Name: "Corso a Scelta II", Credits: 6, Semester: 2, Passed: false},
			},
		},
	}, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// DocumentRepository
// ─────────────────────────────────────────────────────────────────────────────

type DocumentRepository struct{}

func NewDocumentRepository() *DocumentRepository { return &DocumentRepository{} }

func (r *DocumentRepository) GetDocuments(_ context.Context) ([]domain.Document, error) {
	return []domain.Document{
		{Name: "Certificato di iscrizione", Type: "PDF", Size: "245 KB", Date: "15 ott 2023", Status: "Verified"},
		{Name: "Piano di studi 2023/24", Type: "PDF", Size: "128 KB", Date: "5 set 2023", Status: "Verified"},
		{Name: "Documento d'identità", Type: "JPG", Size: "1.2 MB", Date: "20 ago 2023", Status: "Pending"},
		{Name: "Certificato Inglese B2", Type: "PDF", Size: "340 KB", Date: "12 giu 2023", Status: "Verified"},
		{Name: "Ricevuta tasse A.A. 2023/24", Type: "PDF", Size: "89 KB", Date: "30 set 2023", Status: "Verified"},
	}, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// CalendarRepository
// ─────────────────────────────────────────────────────────────────────────────

type CalendarRepository struct{}

func NewCalendarRepository() *CalendarRepository { return &CalendarRepository{} }

func (r *CalendarRepository) GetEvents(_ context.Context) ([]domain.CalendarEvent, error) {
	return []domain.CalendarEvent{
		{Course: "Sistemi Operativi", Type: domain.EventTypeExam, Date: "May 28, 2024", Time: "09:00", Room: "Aula A1", Urgent: true},
		{Course: "Ingegneria del Software", Type: domain.EventTypeDeadline, Date: "Jun 5, 2024", Time: "23:59", Urgent: true},
		{Course: "Intelligenza Artificiale", Type: domain.EventTypeExam, Date: "Jun 15, 2024", Time: "10:30", Room: "Aula B3", Urgent: false},
		{Course: "Statistica", Type: domain.EventTypeExam, Date: "Jul 3, 2024", Time: "09:00", Room: "Aula C2", Urgent: false},
		{Course: "Tesi di Laurea I", Type: domain.EventTypeDeadline, Date: "Sep 30, 2024", Time: "12:00", Urgent: false},
	}, nil
}
