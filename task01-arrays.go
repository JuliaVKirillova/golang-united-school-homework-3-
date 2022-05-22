package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0
	var counter float32 = 0
	for _, val := range input {
		sum = sum + val
		counter++
	}
	return sum / counter
}
