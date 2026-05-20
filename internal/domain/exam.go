package domain

// Exam represents a passed exam in the student's academic career.
type Exam struct {
	Course  string `json:"course"`
	Area    string `json:"area,omitempty"`
	Date    string `json:"date"`
	Grade   int    `json:"grade"`
	Lode    bool   `json:"lode"`
	Credits int    `json:"credits"`
	Year    int    `json:"year,omitempty"`
}

// UpcomingExam represents an exam the student is yet to sit.
type UpcomingExam struct {
	Course  string `json:"course"`
	Date    string `json:"date"`
	Credits int    `json:"credits"`
	Urgent  bool   `json:"urgent"`
}
