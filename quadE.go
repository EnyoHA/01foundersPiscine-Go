package piscine

import (
	"fmt"
	"strings"
)

func QuadE(x, y int) {
	if x > 0 && y > 0 {
		fmt.Print("A")
		if x > 1 {
			xAxis := strings.Repeat("B", x-2)
			fmt.Print(xAxis)
			fmt.Print("C")
		}
		fmt.Print("\n")
		index := 0
		for index < y-2 {
			fmt.Print("B")
			if x > 1 {
				yAxis := strings.Repeat(" ", x-2)
				fmt.Print(yAxis)
				fmt.Print("B")
			}
			fmt.Print("\n")
			index++
		}
		if y > 1 {
			fmt.Print("C")
			if x > 1 {
				fmt.Print(strings.Repeat("B", x-2))
				fmt.Print("A")
			}
			fmt.Print("\n")
		}
	}
}
