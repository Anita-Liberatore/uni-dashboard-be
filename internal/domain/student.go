package domain

// Student is the root entity of the student bounded context.
type Student struct {
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Area           string `json:"area"`
	StudentID      string `json:"studentId"`
	Email          string `json:"email"`
	PEC            string `json:"pec"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	BirthDate      string `json:"birthDate"`
	BirthPlace     string `json:"birthPlace"`
	TaxCode        string `json:"taxCode"`
	Status         string `json:"status"`
	Year           int    `json:"year"`
	Semester       int    `json:"semester"`
	EnrolledSince  string `json:"enrolledSince"`
	GraduationDate string `json:"graduationDate"`
	Advisor        string `json:"advisor"`
}

// AcademicRecord is a value object summarising the student's academic standing.
type AcademicRecord struct {
	Program   string       `json:"program"`
	Degree    string       `json:"degree"`
	Faculty   string       `json:"faculty"`
	Credits   ProgressPair `json:"credits"`
	Courses   ProgressPair `json:"courses"`
	Electives ProgressPair `json:"electives"`
	GPA       float64      `json:"gpa"`
}

// ProgressPair is a value object for current / total counters.
type ProgressPair struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}
