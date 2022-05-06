package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var slice []int
	for k := range input {
		slice = append(slice, k)
	}
	sort.Ints(slice)
	for _, v := range slice {
		result = append(result, input[v])
	}
	return
}
