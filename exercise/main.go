package main

import (
	"fmt"
	"goroutine-course/exercise/student"
)

func ProcessStudentData(data student.Student, ch chan student.Student, chErr chan error) {
	if data == (student.Student{}) {
		chErr <- student.ErrStudentEmpty
		return
	}

	if data.Name == "" {
		chErr <- student.ErrStudentNotValid
		return
	}

	if data.IPK < 0 {
		chErr <- student.ErrStudentNotValid
		return
	}

	if data.SKS < 0 {
		chErr <- student.ErrStudentNotValid
		return
	}

	data.SetFirstName()
	data.SetLastName()
	data.SetIsPassed()
	data.SetHonors()

	ch <- data
}

func runSequentially() {
	dummyData := []student.Student{
		{Name: "Budi Santoso", IPK: 3.5, SKS: 100},
		{Name: "Joko Susilo", IPK: 3.0, SKS: 100},
		{Name: "Andi Wijaya", IPK: 2.5, SKS: 100},
		{Name: "Dewi Lestari", IPK: 3.75, SKS: 100},
	}

	for i := range dummyData {
		dummyData[i].SetFirstName()
		dummyData[i].SetLastName()
		dummyData[i].SetIsPassed()
		dummyData[i].SetHonors()
	}

	for _, data := range dummyData {
		fmt.Println("Name:", data.Name)
		fmt.Println("First Name:", data.FirstName)
		fmt.Println("Last Name:", data.LastName)
		fmt.Println("IPK:", data.IPK)
		fmt.Println("SKS:", data.SKS)
		fmt.Println("Is Passed:", data.IsPassed)
		fmt.Println("Honor:", data.Honor)
		fmt.Println()
	}
}

func runConcurrently() {
	dummyData := []student.Student{
		{Name: "Budi Santoso", IPK: 3.5, SKS: 100},
		{Name: "Joko Susilo", IPK: 3.0, SKS: -1},
		{Name: "Andi Wijaya", IPK: 2.5, SKS: 100},
		{Name: "Dewi Lestari", IPK: 3.75, SKS: 100},
	}

	ch := make(chan student.Student)
	errCh := make(chan error)

	for _, data := range dummyData {
		go ProcessStudentData(data, ch, errCh)
	}

	for i := 0; i < len(dummyData); i++ {
		// materi (select) : https://dasarpemrogramangolang.novalagung.com/A-channel-select.html
		select {
		case data := <-ch:
			fmt.Println("Name:", data.Name)
			fmt.Println("First Name:", data.FirstName)
			fmt.Println("Last Name:", data.LastName)
			fmt.Println("IPK:", data.IPK)
			fmt.Println("SKS:", data.SKS)
			fmt.Println("Is Passed:", data.IsPassed)
			fmt.Println("Honor:", data.Honor)
			fmt.Println()
		case err := <-errCh:
			fmt.Println(err.Error())
		}
	}
}

// main function concurrently process student data
func main() {
	runConcurrently()
}
