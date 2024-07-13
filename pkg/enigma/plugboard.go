package enigma

import (
	"fmt"
)

type Plugboard struct {
	mappings map[byte]byte
}

func (plugboard *Plugboard) PassCharacter(character byte) byte {
	output, ok := plugboard.mappings[character]
	if !ok {
		return character
	}
	return output
}

// NewPlugboard function returns a new plugboard with a given mapping
// containing a number of letter pairs. Valid mapping item contains exactly
// two capital letters from A to Z. Any other input will cause this
// function to fail
func NewPlugboard(mappings []string) (Plugboard, error) {
	plugboardMappings := make(map[byte]byte)
	for _, mapping := range mappings {
		if len(mapping) != 2 {
			return Plugboard{}, fmt.Errorf("incorrect plugboard mapping %s", mapping)
		}
		input := mapping[0]
		output := mapping[1]
		if input < 'A' || input > 'Z' || output < 'A' || output > 'Z' {
			return Plugboard{}, fmt.Errorf("mapping %s contains invalid characters", mapping)
		}
		_, ok := plugboardMappings[input]
		if ok {
			return Plugboard{}, fmt.Errorf("mapping %s already added to plugboard", mapping)
		}
		_, ok = plugboardMappings[output]
		if ok {
			return Plugboard{}, fmt.Errorf("mapping %s already added to plugboard", mapping)
		}
		plugboardMappings[input] = output
		plugboardMappings[output] = input
	}
	return Plugboard{plugboardMappings}, nil
}
