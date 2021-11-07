package main

import (
	"fmt"
	"strings"
)

func Quad(x, y int) {

	// wont print at at all if either value is less than 0
	for x > 0 && y > 0 {
		// as soon as either value is higher than 0
		// immediatly print a character there
		// in this example A
		// 	THIS IS THE RESULT OF THE FIRST FOR STATEMENT
		fmt.Print("A")
		// then go into the repeatable part of the middle
		// THIS IS THE FIRST IF STATEMENT
		if x > 1 {
			// this x > 1 leaves no arguments with the previous print
			// use the strings.Repeat function here which repeats a
			// character for the defined amount of times.
			// in this example x - 2 to leave a space for the first
			// and the last character
			xAxis1 := strings.Repeat("-", x-2)
			// then print this variable followed by the final character
			fmt.Print(xAxis1)
			fmt.Print("B")

		}
		// then end the line by using \n
		fmt.Print("\n")
		// here we set up the index to make sure the right thing
		//is printed on the yAxis
		// could be called yAxis really? may make better sense to me?
		//uses the y variable to know what to print
		yAxis := 0

		// here we do a similar thing to xAxis, where we
		// make sure to leave space for the first line
		// in the first for and if statements.
		for yAxis < y-2 {
			// this print statement makes sure the character provided
			//prints as the first character on each line following the first
			// like along the yAxis of a graph
			fmt.Print("|")
			// then this function moves on to print the characters on
			// the xAxis of line 2. it is a repeat of the previous
			// xAxis function but with different inputs.
			if x > 1 {
				xAxis2 := strings.Repeat(" ", x-2)
				fmt.Print(xAxis2)
				fmt.Print("|")

			}
			// this Print statement then ends the line
			fmt.Print("\n")
			// this statement is what makes this repeat until
			// it reaches y > 2 similar to the xAxis
			// variable and the strings.Repeat package.
			// could be a way to do this without using "strings"?
			yAxis++
		}

		// y has to be greater than 1 for the bottom line
		// to be called in
		if y > 1 {
			fmt.Print("C")
			// for the next if function, you don't have to create
			// a variable as you have the others but
			// could be good practice?
			//makes it clearer for me
			if x > 1 {
				xAxis3 := strings.Repeat("-", x-2)
				fmt.Print(xAxis3)
				fmt.Print("D")
			}
			//end the line
			fmt.Print("\n")
		}

	}
}

func main() {

	Quad(3, 7)
}
