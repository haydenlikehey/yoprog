package main

import (
	"fmt"

	"yoprog/functions" // Note to self: Name the root the name of the project
)
func main() {

	fmt.Printf("Yo, what's good world?\n")

	functions.ReturnFunc()()
	functions.PrintResult(functions.MultipleReturns, 10, 3)

	// See, the below wouldn't work
	// printResult(returnFunc, 3, 10)
	// printResult() doesn't accept the same parameters or return the proper types

	fmt.Println(functions.Hello("Nice, right?"))
	fmt.Println(functions.MultipleReturns(10, 3))
	fmt.Println(functions.MultipleReturns2(10, 3))
}
