package homework

import (
	"sort"
)
func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	a := []int{}
	for key := range input {
		a = append(a, key)
	}
	sort.Ints(a)
	result = []string{}
	for _, value := range a {
		result = append(result, input[value])
	}
	return result
}
