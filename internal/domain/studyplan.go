package domain

// YearPlan groups all courses for a single academic year.
type YearPlan struct {
	Year    int      `json:"year"`
	Courses []Course `json:"courses"`
}

// Course is a value object within the study plan.
// Grade is a pointer so it is omitted from JSON when the exam has not been passed.
type Course struct {
	Name     string `json:"name"`
	Credits  int    `json:"credits"`
	Semester int    `json:"semester"`
	Passed   bool   `json:"passed"`
	Grade    *int   `json:"grade,omitempty"`
	Lode     bool   `json:"lode,omitempty"`
}
