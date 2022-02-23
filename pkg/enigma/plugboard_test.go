package enigma

import (
	"testing"
)

func TestPlugboard_PassCharacter(t *testing.T) {
	// case 1: incorrect mapping length
	{
		_, err := NewPlugboard([]string{"AB", "ERROR"})
		if err == nil {
			t.Errorf("Error in NewPlugboard function")
		}
	}
	// case 2: duplicated mappings
	{
		_, err := NewPlugboard([]string{"AB", "CD", "BD"})
		if err == nil {
			t.Error("Error in NewPlugboard function")
		}
	}

	// case 3: incorrect symbols in mapping
	{
		_, err := NewPlugboard([]string{"AB", "12"})
		if err == nil {
			t.Error("Error in NewPlugboard function")
		}
	}

	// case 4: finally a valid mapping
	{
		plugboard, err := NewPlugboard([]string{"AD", "EG", "ZP"})
		if err != nil {
			t.Errorf("error in NewPlugboard function: %v", err)
		}
		result := plugboard.PassCharacter('A')
		expected := byte('D')
		if result != expected {
			t.Errorf("error in plugboard: expected %c, got %c", result, expected)
		}

		result = plugboard.PassCharacter('G')
		expected = byte('E')
		if result != expected {
			t.Errorf("error in plugboard: expected %c, got %c", result, expected)
		}

		result = plugboard.PassCharacter('B')
		expected = byte('B')
		if result != expected {
			t.Errorf("error in plugboard: expected %c, got %c", result, expected)
		}

		result = plugboard.PassCharacter('C')
		expected = byte('C')
		if result != expected {
			t.Errorf("error in plugboard: expected %c, got %c", result, expected)
		}
	}
}
