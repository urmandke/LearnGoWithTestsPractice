package arrays

//Sum takes an array of 5 integers and returns thier sum as int
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}
