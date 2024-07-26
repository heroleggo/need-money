package main

import (
	"gogogo/exercise"
	"gogogo/utils"
)

func main() {
	exercise.SayHello()
	utils.Parse("<?xml asdasd ?><test><hehe>hoho</hehe></test>", "xml")
}
