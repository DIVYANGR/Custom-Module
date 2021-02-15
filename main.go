package main

import (
	"fmt"

	"example.com/Maths"
)

func main() {
	// Get a greeting message and print it.
	message := Maths.Square(64)
	fmt.Println(message)

	var a, b float64
	fmt.Print("Enter the values for addition  \n")
	fmt.Scanln(&a, &b)
	message1 := Maths.Add(a, b)
	fmt.Println(message1)

	var c, d float64
	fmt.Print("Enter the values for multiplication  \n")
	fmt.Scanln(&c, &d)
	message2 := Maths.Mul(c, d)
	fmt.Println(message2)

	var r float64
	fmt.Print("Enter the values for redius  \n")
	fmt.Scanln(&r)
	message3 := Maths.CircleArea(r)
	fmt.Println(message3)
}
