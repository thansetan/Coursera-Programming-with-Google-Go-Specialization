package main

import "fmt"

/*
Write a program which prompts the user to enter a floating point number
and prints the integer which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.
*/

func main() {
	var floatInput float64
	fmt.Printf("input a float number: ")
	fmt.Scan(&floatInput)
	fmt.Printf("here's your truncated number: %d\n", int(floatInput))
}
