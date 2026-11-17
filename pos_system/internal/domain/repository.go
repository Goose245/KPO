package domain

type StudentRepository interface {
	Save(student *Student) error
	FindByID(id string) (*Student, error)
	FindAll() ([]*Student, error)
}

type CriterionRepository interface {
	Save(criterion *Criterion) error
	FindByID(id string) (*Criterion, error)
	FindAll() ([]*Criterion, error)
}

type EvaluationRepository interface {
	Save(eval *Evaluation) error
	FindByStudentID(studentID string) ([]Evaluation, error)
	FindAll() ([]Evaluation, error)
	Delete(studentID, criterionID string) error
}
