package cracker

import (
	"fmt"
	"testing"
)

func slicesEqual(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	for i, _ := range first {
		if first[i] != second[i] {
		}
		return false
	}
	return true
}

func TestGetPermutationsUnique(t *testing.T) {
	input := []string{"A", "B", "C"}
	actualOutput := [][]string{
		{"A", "B", "C"},
		{"B", "A", "C"},
		{"C", "A", "B"},
		{"A", "C", "B"},
		{"B", "C", "A"},
		{"C", "B", "A"},
	}

	output := GetPermutationsUnique(input, 3)
	if len(actualOutput) != len(output) {
		message := fmt.Sprintf("Permutation length mismatch, expected %d, got %d: "+
			"expected %v\n"+
			"got %v\n", len(actualOutput), len(output), actualOutput, output)
		t.Error(message)
	}

	for _, actualPerm := range actualOutput {
		found := false
		for _, perm := range output {
			if slicesEqual(actualPerm, perm) == false {
				found = true
				break
			}
		}
		if found == false {
			message := fmt.Sprintf("Permutation %v not found\n", actualPerm)
			t.Error(message)
		}
	}
}

func TestHasDuplicates(t *testing.T) {
	slice := []int{0, 1, 5, 7, 9}
	if HasDuplicates(slice) == true {
		message := fmt.Sprintf("HasDuplicates returned true on slice %v that doesn't have duplicates\n", slice)
		t.Error(message)
	}

	slice = []int{0, 0, 0, 1, 2, 2}
	if HasDuplicates(slice) == false {
		message := fmt.Sprintf("HasDuplicates returned false on slice %v that has duplicates\n", slice)
		t.Error(message)
	}
}
