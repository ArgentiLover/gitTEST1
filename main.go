package main

import (
	"fmt"
)

func main() {
	fmt.Println(razn(summm(2, 5, 8), 6))
}

func summm(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func summmm(num1, num2, num3, num4 int) int {
	return num1 + num2 + num3 + num4
}

func razn(num1, num2 int) int {
	return num1 - num2
}
