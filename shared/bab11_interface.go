package shared

import (
	"fmt"
	"strings"
)

func ContohInterface() {
	var secret interface{}
	var data map[string]interface{}
	data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}

	fmt.Println(data)
	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	var secret3 interface{}

	secret3 = 2
	var number = secret3.(int) * 10
	fmt.Println(secret3, "multiplied by 10 is :", number)

	secret3 = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret3.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")

	type person struct {
		name string
		age  int
	}

	var secret4 interface{} = &person{name: "imron", age: 17}
	var name = secret4.(*person).name
	fmt.Println(name)

	// var person = []map[string]interface{}{
	// 	{"name": "Wick", "age": 23},
	// 	{"name": "Ethan", "age": 23},
	// 	{"name": "Bourne", "age": 22},
	// }

	// for _, each := range person {
	// 	fmt.Println(each["name"], "age is", each["age"])
	// }

	var person2 = []map[string]interface{}{
		{"name": "imron", "age": 23},
		{"name": "budi", "age": 25},
		{"name": "ahmad", "age": 27},
	}
	fmt.Println(person2)
	for _, each := range person2 {
		fmt.Println(each["name"], "age is", each["age"])
	}

	var fruits3 = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, each := range fruits3 {
		fmt.Println(each)
	}
}
