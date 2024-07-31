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

const (
	Red   int = iota
	Green int = iota
	Blue  int = iota
)

func PrintColor() {
	fmt.Println(Red)
	fmt.Println(Green)
	fmt.Println(Blue)
}

const typelessConst = 1

func FuncWithForLoop() {
	for {
		if typelessConst == 1 {
			fmt.Println("loop break 1")
			break
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func FunWithIf() {
	if typelessConst == 1 {
		fmt.Println("if statement")
	} else {
		fmt.Println("else statement")
	}
}

func FunWithSwitch() {
	switch typelessConst {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("default case")
	}
}

func HandleArray() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
