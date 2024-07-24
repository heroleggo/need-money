package main

import (
	"fmt"
	"gogogo/exercise"
	"gogogo/utils"
)

func main() {
	exercise.SayHello()
	utils.Parse("test", "xml")

	txt := "hehe"
	txt2 := "가나다라마바사"
	txt3 := '가'

	fmt.Println(txt)
	fmt.Printf("%T\n", txt)
	fmt.Println(txt2)
	fmt.Printf("%T\n", txt2)
	fmt.Println(txt3)
	fmt.Printf("%T\n", txt3)
}
