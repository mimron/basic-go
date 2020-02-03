package shared

import "fmt"

func ContohFunction() {
	addNumbers(353454, 99999)
	addNumbers(353, 9999)
	addNumbers(3554, 99)

	a := addInts(99, 1)
	fmt.Println("a = ", a)

	a = addInts(9, 675)
	fmt.Println("a = ", a)

	div, remainder := divAndRemainder(57, 7)
	fmt.Println(div, remainder)

	// use underscore to ignore a returned value
	div, _ = divAndRemainder(57, 7)
	fmt.Println(div)

	_, remainder = divAndRemainder(57, 7)
	fmt.Println(remainder)

	divAndRemainder(57, 7)

	// in Go, all functions calls are done by value// (exceptions, see later)
	// a copy of input argument variables is passed o the function
	y := 5
	arr := [2]int{45, 99}
	s := "olo"
	doubleFail(y, arr, s)
	fmt.Println("outside doublefail", y, arr, s)
}

// where function is placed does not matter
// no overloading of function w/ different input parameters
func addNumbers(a int, b int) {
	fmt.Println(a + b)
}

func addInts(c int, d int) int {
	return c + d
}

// multiple returns
func divAndRemainder(e int, f int) (int, int) {
	return e / f, e % f
}

func doubleFail(a int, arr [2]int, s string) {
	a = a * 2
	for index := 0; index < len(arr); index++ {
		arr[index] *= 2
	}
	s = s + s
	fmt.Println("in doublefail", a, arr, s)
}
