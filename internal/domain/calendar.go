package domain

// EventType restricts the allowed values for CalendarEvent.Type.
type EventType string

const (
	EventTypeExam     EventType = "Exam"
	EventTypeDeadline EventType = "Deadline"
	EventTypeLecture  EventType = "Lecture"
)

// CalendarEvent is a scheduled event (exam, deadline, or lecture).
type CalendarEvent struct {
	Course string    `json:"course"`
	Type   EventType `json:"type"`
	Date   string    `json:"date"`
	Time   string    `json:"time"`
	Room   string    `json:"room,omitempty"`
	Urgent bool      `json:"urgent"`
}
