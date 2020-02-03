package shared

import "fmt"

func ContohArray() {
	// zero-value of arrays is an array// of specified length of zero-values of the type inside the array
	var myArrInt [4]int
	fmt.Println(myArrInt)

	myArrInt[0] = 12
	myArrInt[1] = 23
	myArrInt[2] = 34
	myArrInt[3] = 45
	fmt.Println(myArrInt)

	// one-line array assihnement in a composite literal expression
	// the length of the array is part of its type definition
	myArrInt = [4]int{111, 222, 333, 444}
	fmt.Println(myArrInt)

	myArrStr := [4]string{"titi", "tooi", "tatu", "teti"}
	fmt.Println(myArrStr)

	// iterate over an array
	for i, val := range myArrInt {
		fmt.Printf("At index %d = %d \n", i, val)
	}

	// slice of an array
	myArrStr2 := myArrStr[:2]

	// use _ to ignore the index variable
	for _, val := range myArrStr2 {
		fmt.Println(val)
	}
	// use of arrays is limited, slices are more flexible

}
