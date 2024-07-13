package cracker

import "gonum.org/v1/gonum/stat/combin"

func GeneratePermutations(n, length int, sender chan<- []int) {
	defer close(sender)
	generator := combin.NewPermutationGenerator(n, length)
	for {
		if !generator.Next() {
			break
		}
		perm := generator.Permutation(nil)
		sender <- perm
	}
}
