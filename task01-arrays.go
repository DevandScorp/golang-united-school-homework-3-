package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var arrSum float32 = 0.0
	for _, value := range input {
		arrSum += value
	}
	return arrSum / float32(len(input))
}
