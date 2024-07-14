package functions

import "fmt"

// Parentheses can be omitted if there is only one parameter
func Hello(name string) string {
	return "This is being returned with no parentheses. " + name
}

// Literally so happy about this feature even though I used to be a structure return advocate!
func MultipleReturns(a int64, b int64) (int64, int64) {
	return a + b, a - b
}

func MultipleReturns2(a int64, b int64) (c int64, d int64) {
	x, y := a+b, a-b
	c = x
	d = y
	// Return names don't need to be specified if using named return values in func definition
	return
}

// Passing a function as a parameter
// Note that I think this is talking about a very specific function.
// f needs to accepts 2 int64's and return 2 int64's to work.
func PrintResult(f func(int64, int64) (int64, int64), a int64, b int64) {
	c, d := f(a, b)
	fmt.Println(c, d)
}

// Functions can return functions as well!
func ReturnFunc() func() {
	return func() {
		fmt.Println("This is a function returned by another function.")
	}
}