package shared

import "fmt"

func ContohString() {
	// zero-value for strings is an empty string ""
	var str string
	fmt.Println("\"" + str + "\"")

	str = "This is a string."
	fmt.Println(str)

	str2 := "Another string."
	fmt.Println(str2)

	// raw string literalss with back ticks
	// can be written on multiple lines and no escapes
	str3 := `
        Raw string in the building.
        And another line.
    `
	fmt.Println(str3)

	str4 := "Dunya nzuri = "
	str5 := "美麗的世界"
	str6 := str4 + str5
	fmt.Println(str6)

	// In Go, strings are immutable sequences of bytes
	// you can access each byte
	str7 := "mazoezi"
	b1 := str7[0]
	b2 := str7[1]
	fmt.Println(str7, "\n\t", b1, "=", string(b1), b2, "=", string(b2))

	// substrings
	s1 := str7[0:2]
	s2 := str7[2:4]
	s3 := str7[:3]
	s4 := str7[3:]
	fmt.Printf("%s \t %s \t %s \t %s \n", s1, s2, s3, s4)

	// length of string
	fmt.Println(str7, " = ", len(str7), " characters")

	// single character = rune -> numeric type, sane as int32
	// can be converted to a string
	var r rune
	// single qutoes
	r = '✖' // same as r = 10006
	fmt.Println("This is a rune : ", r, " which in string = ", string(r))
}
