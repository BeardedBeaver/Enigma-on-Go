package cracker

import "gonum.org/v1/gonum/stat/combin"

func GeneratePermutations(n, length int) [][]int {
	permutations := [][]int{}
	generator := combin.NewPermutationGenerator(8, 3)
	for {
		if !generator.Next() {
			break
		}
		permutations = append(permutations, generator.Permutation(nil))
	}
	return permutations
}
