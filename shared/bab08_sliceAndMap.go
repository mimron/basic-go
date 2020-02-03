package shared

import "fmt"

func ContohSlice() {
	/*
	   SLICE = growable sequence of values of a single specified type
	   the size is not part of the type definition
	*/
	// define a slice in a composite literal expression
	myFiboSlice := []int{0, 1, 2, 3, 5, 8, 13}
	fmt.Println("myFiboSlice is", myFiboSlice)

	// create a slice from another slice = slice expression
	myFiboSlice2 := myFiboSlice[1:4]
	fmt.Println("myFiboSlice2 is", myFiboSlice2)

	// carefully with subslices because// they point to the same location in memory as the original slice
	myFiboSlice[2] = 00
	fmt.Println("\n myFiboSlice is", myFiboSlice)
	// slices are reference types, behave like pointers
	fmt.Println("myFiboSlice2 after modifying myFiboSlice is", myFiboSlice2)

	// zero-value for slice is nil slice (no value in slice)
	titiSlice := []string{"titi_one", "titi_two", "titi_three"}
	var totoSlice []string
	fmt.Println("totoSlice is", totoSlice)
	fmt.Println("length of totoSlice is", len(totoSlice))

	// assigning slice to another slice makes them share same location// in memrory
	totoSlice = titiSlice
	fmt.Println("totoSlice is now", totoSlice)
	titiSlice[0] = "titi_zero"
	fmt.Println("after modifying titiSlice, totoSlice is now", totoSlice)

	// this behavior also happens in functions
	modifySlice(titiSlice)
	fmt.Println("after modifying titiSlice in a function, totoSlice is now", totoSlice)

	// define a slice filled with wero-values of the type specified
	// last argument is the capacity of the underlying array
	s1 := make([]int, 5, 20)

	// copy a slice from another slice with the copy built-in function// "copy in slice s1 the elements of slice myFiboSlice
	copy(s1, myFiboSlice)
	fmt.Println("\ns1 is", s1)

	// append built-in function -> returns a new slice
	// increases the length of the slice
	s2 := append(s1, 21, 34, 55)
	fmt.Println("s2 is", s2)

	// append a slice to another slice
	s3 := []int{111, 222, 333}
	// notice the ... after the s3 identier to spread the elements in it
	s3 = append(s2, s3...)
	fmt.Println("s3 is", s3)

	// deleting element at offset 6 (7th) from slice (the easy way)
	s3 = append(s3[:6], s3[7:]...)
	fmt.Println("s3 is", s3)

	// deleting from slice, the more involved way (see the function below)// it was tough to find a working algorithm !
	s3 = deleteItemFromSliceAtIndex(s3, 3)
	fmt.Println("s3 is", s3)

	// slices are based on an underlying array // for which you can specify the capacity // (number of spots to allocate in memory)
	// here the slice is initialized with 5 zero-value spots // but the underlying array has 100 spots in case we append.
	// This allows to avoid to copy and create a new underlying array // every time we append more than the length of the slice
	s4 := make([]int, 5, 100)
	fmt.Println("s4 is", s4)
	fmt.Println("length of s4 is", len(s4))
	fmt.Println("capacity of s4 is", cap(s4))

	// make a slice of bytes out of a string
	hello := "李先生你好"
	myByteSlice := []byte(hello)
	fmt.Println("\nmyByteSlice is", myByteSlice)

	myRuneSlice := []rune(hello)
	fmt.Println("\nmyRuneSlice is", myRuneSlice)

	// multidimensional slice
	s5 := []int{84, 64, 44}
	s6 := []int{42, 32, 22}
	s7 := [][]int{s5, s6}
	fmt.Println("\ns7 is", s7)

	/********** MAPS ****************/

	// associate value of single data type to value of another data type// collection of key / value pairs (key not restricted to a string)
	// maps are unordered

	myMap := make(map[string]string)
	myMap["name"] = "GOTO"
	myMap["firstname"] = "Florian"
	myMap["occupation"] = "Software Engineer"
	myMap["native_language"] = "French"
	fmt.Printf("\n%v\n", myMap)

	// access value in map
	fmt.Printf("The name is %v\n", myMap["name"])

	// if no value on a key return zero-value for the type of value
	fmt.Printf("The age is %v\n", myMap["age"])

	// make sure a key is in the map = comma ok idiom
	// v = value associated w/ existing key
	// ok = boolean is key in map
	if v, ok := myMap["isBillionaire"]; ok {
		fmt.Println("isBillionaire in map =", v)
	} else {
		// ok == false
		fmt.Println("isBillionaire in map =", ok)
	}

	// map literal declaration (composite literal expression)
	worldMap := map[int]string{
		// every line must end with a comma
		1: "nimefurahi kukuona",
		2: "приємно бачити вас",
		3: "तुम्हें देखकर अच्छा लगा",
		4: "ስለተያየን ደስ ብሎኛል",
		5: "ดีใจที่ได้พบคุณ",
	}

	// iterate over a map - order of iteration is random
	for keyInMap, valueInMap := range worldMap {
		fmt.Println(keyInMap, "=", valueInMap)
	}

	// delete value from map (built-in function)
	delete(worldMap, 2)
	fmt.Println("worldMap :", worldMap)

	// same as slices, maps are passed by reference

	// nil map
	var tMap map[string]int
	// writing to zero-valued map will make the program panic
	// tMap["toto"] = 12345678
	fmt.Println("tMap :", tMap)
	fmt.Println("length of tMap :", len(tMap))

	// delete an element in map
	sport := map[string]string{"yoyo": "ok", "ping pong": "great"}
	fmt.Println("sport :", sport)
	delete(sport, "yoyo")
	fmt.Println("sport :", sport)

	// to make sure that you delete an existing pair in the map
	langs := map[string]string{
		"ES6+":         "great",
		"Go":           "cool",
		"TypeScript":   "ok",
		"Python":       "over hyped but nice",
		"Bash":         "necessary",
		"HTML5":        "necessary",
		"CSS":          "necessary",
		"Elm":          "niche",
		"Java":         "no comments...",
		"Rust":         "to assess",
		"Web Assembly": "who knows",
	}
	fmt.Println("lang :", langs)
	if _, ok := langs["assembly"]; ok {
		delete(langs, "assembly")
	}
	fmt.Println("lang :", langs)
}

func modifySlice(s []string) {
	s[1] = "one"
	fmt.Println("\n", s)
}

func deleteItemFromSliceAtIndex(s []int, index int) []int {
	temp := make(map[int]int, len(s))
	// convert slice to map
	for i, v := range s {
		if i == index {
			continue
		}
		temp[i] = v
	}

	newSlice := make([]int, len(temp))

	for i := 0; i < len(temp); i++ {
		if i == index {
			// comma dot idiom = check if key in map
			value, ok := temp[i+1]

			// check that next index exists
			if ok && i <= len(temp) {
				newSlice[i] = value
				continue
			}
		}

		if i > index {
			newSlice[i] = temp[i+1]
			continue
		}

		newSlice[i] = temp[i]
	}

	return newSlice
}
