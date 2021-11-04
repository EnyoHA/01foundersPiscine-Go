package main

import "fmt"

func BasicAtoi(s string) int {

	numbers := []rune(s)
	numbers2 := []int(numbers)


	for i := len(s) - 1; i >= 0; i-- {
		num1 += s[i] * op
		op *= 10
	}
	return res

}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
