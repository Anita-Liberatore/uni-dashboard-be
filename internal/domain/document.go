package domain

// Document represents a file uploaded by the student.
type Document struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Size   string `json:"size"`
	Date   string `json:"date"`
	Status string `json:"status"`
}
