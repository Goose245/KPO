package domain

type Student struct {
	ID   string
	Name string
}

func NewStudent(id, name string) *Student {
	return &Student{
		ID:   id,
		Name: name,
	}
}
