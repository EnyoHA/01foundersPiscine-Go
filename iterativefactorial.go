


func IterativeCalculation(index int) int {

	result := 0

	for i := 0; i < result+1; i++ {
		result = result + 1
	}

	return result

}

func recursiveCalculation(index int) int {

	if index == 1 {
		return 1
	}
	if index > 1 {
		return index + recursiveCalculation(index-1)
	}
	return 0

}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
