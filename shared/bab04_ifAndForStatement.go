package shared

import "fmt"

func ContohIfAndForStatement() {
	a := 27

	// no parentheses surrounding the condition
	// the body of the if statement MUST be surrounded with {} // no matter what
	// no truthy values
	if a > 25 {
		fmt.Println("a is greater than 25")
	} else {
		fmt.Println("a is less than 25")
	}

	b := 546

	// if statement have block scope
	if b == 546 {
		// c does not exist outside of the if
		c := 54
		fmt.Println(b + c)
	}

	// can declare a variable available ONLY in if and else block
	if d := 44; b < 25 {
		fmt.Println("a is greater than 25", d)
	} else {
		fmt.Println("a is less than 25", d)
	}

	// for loop, no parentheses around signature
	e := 2
	for index := 0; index < 10; index++ {
		if index+e == 2 {
			continue
		}
		if index > 8 {
			break
		}
		fmt.Println("index =", index)

	}

	// equivalent of while statement
	f := 0
	for f < 6 {
		fmt.Println("f =", f)
		// don't forget to have smth allowing to get out of the loop
		f++
	}

	// infinite loop
	g := 0
	for {
		fmt.Println("g =", g)
		g++
		// to stop it at some point
		if g > 30 {
			break
		}
	}

	// for range loop

	h := "this is great!"
	for k, v := range h {
		fmt.Println("offset (position) =", k, ", value as rune =", v, ",value as string = ", string(v))
	}

	// logical operators
	fmt.Printf("%t && %t is %t \n", true, false, true && false)
	fmt.Printf("%t || %t is %t \n", true, false, true || false)
	fmt.Printf("!%t is %t \n", true, !true)
}
