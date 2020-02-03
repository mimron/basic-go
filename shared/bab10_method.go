package shared

import "fmt"

func ContohMethod() {
	//contoh method
	var s1 = student{"john wick", 21}
	s1.sayHello()

	//contoh method dengan variable parameter
	var s2 = student{"imron sangat ganteng", 21}
	var a = s2.getNameAt(1)
	fmt.Println("haloss", a)
}
