//Data Types
// string
// bool
// int
// int int8 int32 int64
// uint uint8 uint16 uint64
// byte  alias for uint8
// rune  alias for int32
// float32 float64
// complex64 complex128   for large numbers

package main

import "fmt"

func main2() {
	var name = "GZ"
	apellido := "Zapata" //shorthand for variable
	size := 1.3          //by default is float64
	var age int = 32
	const isMale bool = true
	correo, username := "tavoride@h", "gustavozapata"
	fmt.Println("Hola Mundo", name, apellido, "and", age, correo, username)
	fmt.Printf("%T\n", size) //prints the type of the variable
}
