package shared

import "fmt"

func ContohVariable() {

	var toto int8 = 123
	fmt.Println(toto)

	var tito = 123
	fmt.Println(tito)

	// variable declaration with type inference
	toti := 123
	fmt.Println(toti)

	// variable declaration (implicitly initialized to zero-value,// for numeric types = 0)
	var tata int
	fmt.Println(tata)

	// variable assignment
	tata = 951
	fmt.Println(tata)
}
