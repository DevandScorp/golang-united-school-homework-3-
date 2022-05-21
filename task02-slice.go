package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	reversedSlice := []int64{}
	for i := range input {
		reversedSlice = append(reversedSlice, input[len(input) - i - 1])
	}
	return reversedSlice
}
