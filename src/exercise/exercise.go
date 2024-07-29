package exercise

import "fmt"

func SayHello() {
	fmt.Println("Hello")
}

func functionName(a int, b string) int {
	return a
}

func functionName2(a int, b string) string {
	return b
}

func functionName3(a int, b int) int {
	return a + b
}

func MultiReturnFunc(a int, b int) (int, string) {
	return a + b, "hehe"
}

func FuncWithResultValueName(a int) (result int) {
	result = a
	return
}
