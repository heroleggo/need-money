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

type a struct {
	aa int
	bb string
	cc float64
}

type b struct {
	aa int
	bb string
	cc float64
	a
}

type c struct {
	aa   int
	bb   string
	cc   float64
	data a
}

func UseStruct() {
	a1 := a{1, "hehe", 1.1}
	b1 := b{2, "haha", 2.2, a1}
	c1 := c{3, "hihi", 3.3, a1}

	fmt.Println("print struct b")
	fmt.Println(b1.a.aa)
	fmt.Println(b1.aa)

	fmt.Println("print struct c")
	fmt.Println(c1.data.aa)
	fmt.Println(c1.aa)
}
