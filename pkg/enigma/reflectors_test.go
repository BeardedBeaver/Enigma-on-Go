package enigma

import (
	"fmt"
	"testing"
)

func TestReflectors(t *testing.T) {
	var input = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for model, mapping := range reflectorMappings {
		reflector, err := NewReflector(model)
		fmt.Println("Testing Reflector", model)
		if err != nil {
			t.Error("Can't create Reflector model", model)
		}
		for i, _ := range input {
			result := reflector.Pass(input[i])
			if result != mapping[i] {
				message := fmt.Sprintf("Error on Reflector "+
					model+": "+
					"input %c, "+
					"expected %c, "+
					"got %c",
					input[i], mapping[i], result)
				t.Error(message)
			}
		}
	}
}
