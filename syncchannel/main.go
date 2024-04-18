package main

import "fmt"

type Student struct {
	Name     string
	IPK      float64
	SKS      int
	IsPassed bool
	Honor    string
}

func (s *Student) SetIsPassed(x chan Student) {
	if s.IPK >= 2.5 {
		s.IsPassed = true
	} else {
		s.IsPassed = false
	}

	x <- *s
}

func (s *Student) SetHonors(x, y chan Student) {
	*s = <-x

	if s.IsPassed {
		if s.IPK >= 3.5 && s.SKS >= 100 {
			s.Honor = "Cumlaude"
		} else {
			s.Honor = "No Honor"
		}
	}

	y <- *s
}

func main() {

	listStudent := []Student{
		{Name: "Budi Santoso", IPK: 3.5, SKS: 100},
		{Name: "Joko Susilo", IPK: 3.0, SKS: 100},
		{Name: "Andi Wijaya", IPK: 2.5, SKS: 100},
		{Name: "Dewi Lestari", IPK: 3.75, SKS: 100},
	}

	chStudent := make(chan Student, len(listStudent))
	chStudentHonor := make(chan Student, len(listStudent))

	for _, student := range listStudent {
		// avoid race condition
		student := student
		go func() {
			student.SetIsPassed(chStudent)
			student.SetHonors(chStudent, chStudentHonor)
		}()
	}

	for i := 0; i < len(listStudent); i++ {
		student := <-chStudentHonor
		fmt.Println("Name:", student.Name)
		fmt.Println("IPK:", student.IPK)
		fmt.Println("SKS:", student.SKS)
		fmt.Println("Is Passed:", student.IsPassed)
		fmt.Println("Honor:", student.Honor)
		fmt.Println()
	}

}
