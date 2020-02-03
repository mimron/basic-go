package shared

import (
	"fmt"
	"strings"
)

func ContohStruct() {
	type student struct {
		name  string
		grade int
	}

	var s1 student
	s1.name = "imron"
	s1.grade = 90

	fmt.Printf("The name is %v\n", s1)
	fmt.Println("name  :", s1.name)
	fmt.Println("grade :", s1.grade)

	//cara inisiasi
	var s2 = student{}
	s2.name = "imron2"
	s2.grade = 900

	fmt.Printf("The name is %v\n", s2)
	fmt.Println("name  :", s2.name)
	fmt.Println("grade :", s2.grade)

	var s3 = student{name: "imron3", grade: 100}
	var s4 *student = &s3

	fmt.Println("student 3, name :", s3.name)
	fmt.Println("student 4, name :", s4.name)

	s4.name = "ethan"
	fmt.Println("student 3, name :", s3.name)
	fmt.Println("student 4, name :", s4.name)
}

//=========CONTOH METHOD=========
type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}
