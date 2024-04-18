package main

type Student struct {
	Name     string
	IPK      float64
	SKS      int
	IsPassed bool
	Honor    string
}

func (s *Student) SetIsPassed(x chan Student) {
	// TODO implement code
}

func (s *Student) SetHonors(x, y chan Student) {
	// TODO implement code
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
		go student.SetIsPassed(chStudent)
		go student.SetHonors(chStudent, chStudentHonor)
	}

	for i := 0; i < len(listStudent); i++ {
		student := <-chStudentHonor
		println("Name:", student.Name)
		println("IPK:", student.IPK)
		println("SKS:", student.SKS)
		println("Is Passed:", student.IsPassed)
		println("Honor:", student.Honor)
		println()
	}

}
