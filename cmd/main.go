package main

import (
	"fmt"
	"structs/turma"
)

func main() {
	student1 := turma.Student{
		Name: "Guilherme Moura",
		Age:  22,
	}
	student2 := turma.Student{
		Name: "Luiz Brandi",
		Age:  20,
	}

	fmt.Println(student1)

	course1 := turma.Course{
		Name: "Computer Science",
	}

	course1.Register(student1)
	course1.Register(student2)
	fmt.Println(course1)

	var ptr *turma.Student = &student1
	fmt.Println(ptr)

}
