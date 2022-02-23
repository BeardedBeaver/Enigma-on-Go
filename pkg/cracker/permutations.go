package cracker

func HasDuplicates(slice []int) bool {
	if len(slice) < 1 {
		return false
	}
	var counts = make(map[int]int)
	for _, v := range slice {
		counts[v] += 1
		if counts[v] > 1 {
			return true
		}
	}
	return false
}

func GetPermutations(str []string, amount int) [][]string {
	var result = make([][]string, 0)
	length := len(str)

	// initialize indices with iota-styled stuff
	var indices = make([]int, amount)
	for i, _ := range indices {
		indices[i] = i
	}
	var slice = make([]string, 0)
	for _, index := range indices {
		slice = append(slice, str[index])
	}
	result = append(result, slice)

	for {
		indices[amount-1] += 1
		for i := amount - 1; i > 0; i-- {
			if indices[i] >= length {
				indices[i] = 0
				indices[i-1] += 1
			}
		}
		if indices[0] == length {
			break
		}
		if HasDuplicates(indices) {
			continue
		}
		var slice = make([]string, 0)
		for _, index := range indices {
			slice = append(slice, str[index])
		}
		result = append(result, slice)

	}

	return result
}
