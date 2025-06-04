package hello

import "C"

//export SayHello
func SayHello() string {
	return "Hello from Go!"
}

//export Add
func Add(a, b int) int {
	return a + b
}
