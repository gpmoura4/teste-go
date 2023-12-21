package turma

type Student struct {
	Name string
	Age  uint
}

type Course struct {
	Name     string
	Students []Student
}

// Método pra adicionar novo estudante ao curso
func (c *Course) Register(s Student) {
	c.Students = append(c.Students, s)
}

// Simulando Herança
type EadCourse struct {
	Course  Course
	Website string
}

// Novo tipo de dado float
type NewFlaot float64
