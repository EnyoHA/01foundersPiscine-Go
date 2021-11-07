package main

import (
	"fmt"
	// calling in strings also checks and only prints postive integers, not sure if doing it like this means I can ignore the first part of the if statement?
	"strings"
)

func Quad(x, y int) {
	// this is statement checks that x and y are positive numbers.
	if x > 0 && y > 0 {
		// to print the first character, it checks that both x and y are larger than 0. if they are not it wont run the function.
		fmt.Print("A")
		// if x is greater than 1 it will print xAxis. and the last character on the x axis (set by x).
		if x > 1 {
			// xAxis is made up using the strings repeat package. This works by repeating the "stringed" character the amount of times after the comma.
			// it is x-2 so that the two corner pieces can be different.
			xAxis := strings.Repeat("*", x-2)
			// then prints the following two Print statements in order.
			fmt.Print(xAxis)
			fmt.Print("B")
		}
		// then PRint a \n to end the line
		fmt.Print("\n")
		// here set up an index or _ variable to print the correct things on the yAxis. set as 0.
		index := 0
		// here we then check that when the index is less than y-2 it will print the character furthest left on the yAxis
		for index < y-2 {
			// if it is less than y - 2 it'll print the "character in quotes"
			fmt.Print("*")
			// here we check again if x > 1 and if it is move on to the rest of the function.
			// nuse strings.Repeat again in a new yAxis variable to repeat the " " x-2 times.
			// again to make sure that it leaves space for the beginning and end characters (Its called yAxis despite it also checking the xAxis).
			if x > 1 {
				yAxis := strings.Repeat(" ", x-2)
				// then print this string the required times and ending it with the character in the follwoing print statement
				fmt.Print(yAxis)
				fmt.Print("*")
			}
			// this then ends the line and then will move on to the next by working through the index. the index variable has to be in the for loop
			// im not 100% sure why, but outside it doesnt use it. Works its way down the Y value as it gets higher?
			fmt.Print("\n")
			index++
		}
		// for the bottom line. so y has to be larger than 1 for this line to be called in.
		if y > 1 {
			// the bottom left corner(first character on the line)
			fmt.Print("C")
			// when x is greater than one carry on down the line and just repeats like in the first if loop.
			if x > 1 {
				// to print a single backslash in a Print statement you have to put two in like the example below in strings.Repeat
				fmt.Print(strings.Repeat("\\", x-2))
				fmt.Print("D")
			}
			fmt.Print("\n")
		}
	}
}

// add in different values here to change the size. x first then y like in the Quad function
func main() {
	Quad(5, 10)
}
