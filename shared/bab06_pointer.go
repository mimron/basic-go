package shared

import "fmt"

func ContohPointer() {
	var numberA int = 4
	var numberB *int = &numberA

	var stringA string = "imron"
	var stringB *string = &stringA

	stringA = "egi"
	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	fmt.Println("stringA (value)   :", stringA)  // imron
	fmt.Println("stringA (address) :", &stringA) // 0xc0000703c0

	fmt.Println("stringB (value)   :", *stringB) // imron
	fmt.Println("stringB (address) :", stringB)  // 0xc0000703c0

	//conth parameter pointer
	var number = 4
	fmt.Println("before :", number) // 4

	change(&number, 10)
	fmt.Println("after  :", number) // 10
}

func change(original *int, value int) {
	*original = value
}

func ContohPointer2() {
	a := 10
	// & = reference / pointer to variable "a"
	// the value of b is the location where a is stored
	b := &a
	// c is a copy of "a" at a given time, // they are independent of each other after the first assignment
	c := a
	fmt.Println(a, b, *b, c)

	a = 20
	// to see the value inside the memory location use * // (de-reference the pointer and get to the value)
	fmt.Println(a, b, *b, c)

	// dereference the pointer and assign a value in memory
	// therefore the value of "a" also changes
	*b = 30
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)

	// zero-value for a pointer is nil (absence of value)
	var d *int
	fmt.Println("value of d =", d)
	// cannot read or write value of a nil pointer
	// fmt.Println(*d)   // throws a panic

	e := new(int)
	// new keyword makes a pointer for the type
	fmt.Println("pointer to e =", e)
	// new also allocates memory, here to the zero-value forint type// therefore no panic
	fmt.Println("value of e =", *e)

	f := 20
	fmt.Println("value of f =", f)
	// we pass a pointer to f into that function
	setTo10(&f)
	fmt.Println("value of f after setTo10 =", f)

	g := 30
	fmt.Println("pointer to g =", &g)
	fmt.Println("value of g =", g)
	// on Go, variable in function calls are passed by VALUE
	setTo10Fail(&g)
	fmt.Println("pointer to g after #setTo10Fail =", &g)
	// you CANNOT change the pointer of a variable passed
	fmt.Println("value of g after #setTo10Fail =", g)
}

// pointer as input parameter
func setTo10(pointerToInt *int) {
	*pointerToInt = 10
}

func setTo10Fail(pointerToInt *int) {
	fmt.Println("#setTo!(Fail pointer passed as argument =", pointerToInt)
	// will not affect the original pointer because passed by value
	pointerToInt = new(int)
	fmt.Println("#setTo!(Fail reassignment =", pointerToInt)

	// set the value in memory to 10
	*pointerToInt = 10
}
