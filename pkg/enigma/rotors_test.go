package enigma

import (
	"fmt"
	"testing"
)

func TestRotorsInitialPosition(t *testing.T) {
	var input = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for model, mapping := range rotorMappings {
		rotor, err := NewRotor(model, 'A', 0)
		if err != nil {
			t.Error("Can't create Rotor model", model)
		}
		fmt.Println("Testing Rotor", model)
		for i := range input {
			result := rotor.PassForward(input[i])
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

func TestRotorARotation(t *testing.T) {
	model := "I"
	rotor, err := NewRotor(model, 'A', 0)
	if err != nil {
		t.Error("Can't create Rotor model", model)
	}

	result := rotor.PassForward('A')
	if result != 'E' {
		t.Error("Failed initial position", "expected", 'E', "got", result)
	}

	rotor.Spin()
	result = rotor.PassForward('A')
	if result != 'J' {
		message := fmt.Sprintf("Failed rotated position 1: expected %c, got %c", 'J', result)
		t.Error(message)
	}

	rotor.Spin()
	result = rotor.PassForward('A')
	if result != 'K' {
		message := fmt.Sprintf("Failed rotated position 2: expected %c, got %c", 'K', result)
		t.Error(message)
	}

	rotor.Spin()
	result = rotor.PassForward('A')
	if result != 'C' {
		message := fmt.Sprintf("Failed rotated position 3: expected %c, got %c", 'C', result)
		t.Error(message)
	}

	rotor.Spin()
	result = rotor.PassForward('A')
	if result != 'H' {
		message := fmt.Sprintf("Failed rotated position 4: expected %c, got %c", 'H', result)
		t.Error(message)
	}
}

func TestRotorARotationBackPass(t *testing.T) {
	model := "I"
	rotor, err := NewRotor(model, 'A', 0)
	if err != nil {
		t.Error("Can't create Rotor model", model)
	}

	result := rotor.PassBackward('E')
	if result != 'A' {
		message := fmt.Sprintf("Failed rotated position 1 (back pass): expected %c, got %c", 'A', result)
		t.Error(message)
	}

	rotor.Spin()
	result = rotor.PassBackward('J')
	if result != 'A' {
		message := fmt.Sprintf("Failed rotated position 1 (back pass): expected %c, got %c", 'A', result)
		t.Error(message)
	}

	rotor.Spin()
	result = rotor.PassBackward('K')
	if result != 'A' {
		message := fmt.Sprintf("Failed rotated position 2 (back pass): expected %c, got %c", 'A', result)
		t.Error(message)
	}

	rotor.Spin()
	result = rotor.PassBackward('C')
	if result != 'A' {
		message := fmt.Sprintf("Failed rotated position 3 (back pass): expected %c, got %c", 'A', result)
		t.Error(message)
	}
}

// FIXME this is probably incorrect because an offset here is 0
func TestRotorAWithOffset(t *testing.T) {
	model := "I"
	rotor, err := NewRotor(model, 'B', 0)
	if err != nil {
		t.Error("Can't create Rotor model", model)
	}

	result := rotor.PassForward('A')
	if result != 'J' {
		message := fmt.Sprintf("Failed rotated position 1 (back pass): expected %c, got %c", 'A', result)
		t.Error(message)
	}
}

func TestRotorIIWithOffset(t *testing.T) {
	// case 1
	{
		model := "II"
		rotor, err := NewRotor(model, 'T', 14)
		if err != nil {
			t.Error("Can't create Rotor model", model)
		}

		result := rotor.PassForward('M')
		answer := byte('B')
		if result != answer {
			message := fmt.Sprintf("Failed rotated position 1 (back pass): expected %c, got %c", answer, result)
			t.Error(message)
		}
	}

	// case 2
	{
		model := "II"
		rotor, err := NewRotor(model, 'B', 5)
		if err != nil {
			t.Error("Can't create Rotor model", model)
		}

		result := rotor.PassForward('E')
		answer := byte('E')
		if result != answer {
			message := fmt.Sprintf("Failed rotated position 1 (back pass): expected %c, got %c", answer, result)
			t.Error(message)
		}
	}

	// case 3
	{
		model := "II"
		rotor, err := NewRotor(model, 'W', 4)
		if err != nil {
			t.Error("Can't create Rotor model", model)
		}

		result := rotor.PassForward('F')
		answer := byte('D')
		if result != answer {
			message := fmt.Sprintf("Failed rotated position 1 (back pass): expected %c, got %c", answer, result)
			t.Error(message)
		}
	}
}

func TestNormalizeCharacter(t *testing.T) {
	data := map[int]int{
		0:   0,
		5:   5,
		26:  0,
		29:  3,
		-1:  25,
		-24: 2,
	}

	for key, value := range data {
		result := normalizeCharacter(key)
		if result != value {
			message := fmt.Sprintf("Normalize character error "+
				"input %d, "+
				"expected %d, "+
				"got %d",
				key, value, result)
			t.Error(message)
		}
	}
}

func TestRotor_IsAtNotch(t *testing.T) {
	model := "I"
	rotor, err := NewRotor(model, 'A', 0)
	if err != nil {
		t.Error("Can't create Rotor model", model)
	}

	for rotor.position != 'Q'-'A' { // rotor I as a notch in Q which is hardcoded in this test
		if rotor.IsAtNotch() == true {
			t.Error("Rotor", model, "says it's at notch while in fact it shouldn't be")
		}
		rotor.Spin()
	}

	if rotor.IsAtNotch() == false {
		t.Error("Rotor", model, "says it's not at notch while in fact it should be")
	}
}

func TestRotor_IsAtNotchWithOffset(t *testing.T) {
	model := "I"
	rotor, err := NewRotor(model, 'F', 0)
	if err != nil {
		t.Error("Can't create Rotor model", model)
	}

	for rotor.position != 'Q'-'A' { // rotor I as a notch in Q which is hardcoded in this test
		if rotor.IsAtNotch() == true {
			t.Error("Rotor", model, "says it's at notch while in fact it shouldn't be")
		}
		rotor.Spin()
	}

	if rotor.IsAtNotch() == false {
		t.Error("Rotor", model, "says it's not at notch while in fact it should be")
	}
}
