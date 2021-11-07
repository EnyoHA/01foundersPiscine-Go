package piscine

import (
	"fmt"
	"strings"
)

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		fmt.Print("/")
		if x > 1 {
			xAxis := strings.Repeat("*", x-2)
			fmt.Print(xAxis)
			fmt.Print("\\")
		}
		fmt.Print("\n")
		index := 0
		for index < y-2 {
			fmt.Print("*")
			if x > 1 {
				yAxis := strings.Repeat(" ", x-2)
				fmt.Print(yAxis)
				fmt.Print("*")
			}
			fmt.Print("\n")
			index++
		}
		if y > 1 {
			fmt.Print("\\")
			if x > 1 {
				fmt.Print(strings.Repeat("*", x-2))
				fmt.Print("/")
			}
			fmt.Print("\n")
		}
	}
}
