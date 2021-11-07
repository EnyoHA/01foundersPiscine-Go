package piscine

import "fmt"

func QuadA(x, y int) {
	for yA := 0; yA < y; yA++ {
		for xA := 0; xA < x; xA++ {
			yAxis := (yA+1) == y || yA == 0
			xAxis := (xA+1) == x || xA == 0
			if yAxis && xAxis {
				fmt.Print("o")
			} else if yAxis {
				fmt.Print("-")
			} else if xAxis {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
