package student

import (
	"errors"
	"strings"
)

type Student struct {
	Name      string
	IPK       float64
	SKS       int
	FirstName string
	LastName  string
	IsPassed  bool
	Honor     string
}

var ErrStudentNotValid error = errors.New("student data is not valid")
var ErrStudentEmpty error = errors.New("student name is empty")

func (s *Student) SetHonors() {
	if s.IPK >= 3.5 && s.SKS >= 100 {
		s.Honor = "Cumlaude"
	} else if s.IPK >= 3.5 {
		s.Honor = "Summa Cumlaude"
	} else if s.IPK >= 3.0 {
		s.Honor = "Magna Cumlaude"
	} else {
		s.Honor = "No Honor"
	}
}

func (s *Student) SetFirstName() {
	listName := strings.Split(s.Name, " ")
	s.FirstName = listName[0]
}

func (s *Student) SetLastName() {
	listName := strings.Split(s.Name, " ")
	s.LastName = listName[len(listName)-1]
}

func (s *Student) SetIsPassed() {
	if s.IPK >= 2.75 && s.SKS >= 100 {
		s.IsPassed = true
	} else {
		s.IsPassed = false
	}
}
