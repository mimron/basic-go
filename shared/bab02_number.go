package shared

import "fmt"

func ContohTypeNumber() {

	// Go does not convert types automatically
	// need to explicitly convert them
	var (
		i int8    = 20
		f float32 = 5.6
	)
	fmt.Println(i + int8(f+1.9))

	var (
		j int32 = 456
		k int64 = 987654
	)
	fmt.Println(int64(j) + k)

	// byte is an alias for uint8
	// no need to convert uint8 to byte because same
	var (
		l byte  = 123
		m uint8 = 45
	)
	fmt.Println(l + m)

	// int is an alias for int32 or int64, // dedpending on your mqchine's integer value
	var (
		n int32 = 324
		o int   = 84529899
	)
	fmt.Println(int(n) + o)

	// uint is an alias for uint32 or int64
	var (
		p uint   = 999
		q uint64 = 9999
	)
	fmt.Println(p + uint(q))

	// float operations do not produce an exact result after n decimals// like in most languages
	myFloat := 1.000
	myFloat2 := .999
	fmt.Println(myFloat - myFloat2)

	// arithmetic operations
	fmt.Printf("%d + %d = %d \n", 25, 39, 25+39)
	fmt.Printf("%d - %d = %d \n", 25, 39, 25-39)
	fmt.Printf("%d * %d = %d \n", 25, 39, 25*39)
	fmt.Printf("%d / %d = %v \n", 25, 39, 25/39)
	fmt.Printf("%d %% %d = %v \n", 25, 39, 25%39)

	// constants
	const goldenRatio float64 = 1.6180327868852
	fmt.Printf("The golden ration approximately %f \n", goldenRatio)
	fmt.Printf("The ration truncated to the 3rd decimal is %.3f \n", goldenRatio)

	// formatted printing for numeric types
	fmt.Printf("decimal is %d \n", 99)
	fmt.Printf("binary is %b \n", 99)
	fmt.Printf("unicode reference is %c \n", 99)
	fmt.Printf("hexadecimal is %x \n", 99)
	fmt.Printf("scientific notation of goldenRatio is %e \n", goldenRatio)
}
