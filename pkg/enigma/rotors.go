package enigma

import (
	"errors"
	"fmt"
)

var rotorMappings = map[string]string{
	"I":    "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	"II":   "AJDKSIRUXBLHWTMCQGZNPYFVOE",
	"III":  "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	"IV":   "ESOVPZJAYQUIRHXLNFTGKDCMWB",
	"V":    "VZBRGITYUPSDNHLXAWMJQOFECK",
	"VI":   "JPGVOUMFYQBENHZRDKASXLICTW",
	"VII":  "NZJHGRCXMYSWBOUFAIVLPEKQDT",
	"VIII": "FKQHTLXOCBJSPDZRAMEWNIUYGV",
}

var rotorNotchPositions = map[string][]string{
	"I":    {"Q"},
	"II":   {"E"},
	"III":  {"V"},
	"IV":   {"J"},
	"V":    {"Z"},
	"VI":   {"Z", "M"},
	"VII":  {"Z", "M"},
	"VIII": {"Z", "M"},
}

type Rotor struct {
	mapping        string
	backMapping    string
	position       int
	offset         int
	notchPositions []string
}

func (rotor *Rotor) passCharacter(mapping string, character byte) byte {
	code := int(character - 'A')
	code += rotor.position
	code -= rotor.offset
	code = normalizeCharacter(code)

	result := int(mapping[code] - 'A')
	result -= rotor.position
	result += rotor.offset
	result = normalizeCharacter(result) + 'A'

	return byte(result)
}

// PassForward function passes a signal through the rotor
// in a forward direction without any rotation
func (rotor *Rotor) PassForward(character byte) byte {
	return rotor.passCharacter(rotor.mapping, character)
}

// PassBackward function passes a signal through the rotor
// in a backward direction without any rotation
func (rotor *Rotor) PassBackward(character byte) byte {
	return rotor.passCharacter(rotor.backMapping, character)
}

// Spin spins a rotor by one step changing its position
func (rotor *Rotor) Spin() {
	rotor.position++
	if rotor.position >= 26 {
		rotor.position = 0
	}
}

func (rotor *Rotor) IsAtNotch() bool {
	var currentPosition = fmt.Sprintf("%c", rotor.position+'A')
	for _, position := range rotor.notchPositions {
		if position == currentPosition {
			return true
		}
	}
	return false
}

func normalizeCharacter(character int) int {
	for character < 0 {
		character += 26
	}
	return character % 26
}

// NewRotor function creates a new rotor with a given position (expressed as a capital letter)
// and an offset (expressed as a 0-based index)
func NewRotor(model string, position, offset int) (Rotor, error) {
	if position < 'A' || position > 'Z' {
		return Rotor{}, errors.New("rotor position should be a capital letter A-Z")
	}
	if offset < 0 || offset >= 26 {
		return Rotor{}, errors.New("rotor offset should be a number between 0 and 25")
	}
	mapping, ok := rotorMappings[model]
	if ok == false {
		return Rotor{}, errors.New(fmt.Sprintf("Mapping not found for rotor model %s", model))
	}

	// forms a back mapping for the back pass
	var backMapping string
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, letter := range alphabet {
		for j, letterInMapping := range mapping {
			if letter == letterInMapping {
				backMapping += string(alphabet[j])
			}
		}
	}

	notchPositions, ok := rotorNotchPositions[model]
	if ok == false {
		return Rotor{}, errors.New(fmt.Sprintf("Notch position not found for rotor model %s", model))
	}
	rotor := Rotor{mapping, backMapping, position - 'A', offset, notchPositions}
	return rotor, nil
}
