package shared

import (
	"fmt"
	"reflect"
)

func ContohReflect() {

	var number = 23
	var refelctValue = reflect.ValueOf(number)
	fmt.Println("tipe  variabel :", refelctValue.Type())
	fmt.Println("nilai  variabel :", refelctValue.Interface())

	///Pengaksesan Informasi Property Variabel Objek

}
