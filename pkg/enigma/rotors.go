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

var rotorBackMappings = map[string]string{
	"I":    "UWYGADFPVZBECKMTHXSLRINQOJ",
	"II":   "AJPCZWRLFBDKOTYUQGENHXMIVS",
	"III":  "TAGBPCSDQEUFVNZHYIXJWLRKOM",
	"IV":   "HZWVARTNLGUPXQCEJMBSKDYOIF",
	"V":    "QCYLXWENFTZOSMVJUDKGIARPHB",
	"VI":   "SKXQLHCNWARVGMEBJPTYFDZUIO",
	"VII":  "QMGYVPEDRCWTIANUXFKZOSLHJB",
	"VIII": "QJINSAYDVKBFRUHMCPLEWZTGXO",
}

var rotorNotchPositions = map[string][]byte{
	"I":    {'Q'},
	"II":   {'E'},
	"III":  {'V'},
	"IV":   {'J'},
	"V":    {'Z'},
	"VI":   {'Z', 'M'},
	"VII":  {'Z', 'M'},
	"VIII": {'Z', 'M'},
}

func AvailableRotorModels() []string {
	return []string{
		"I",
		"II",
		"III",
		"IV",
		"V",
		"VI",
		"VII",
		"VIII",
	}
}

type Rotor struct {
	mapping        string
	backMapping    string
	position       int
	offset         int
	notchPositions []byte
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
	var currentPosition = byte(rotor.position + 'A')
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
	if !ok {
		return Rotor{}, fmt.Errorf("mapping not found for rotor model %s", model)
	}

	backMapping, ok := rotorBackMappings[model]
	if !ok {
		return Rotor{}, fmt.Errorf("backwards mapping not found for rotor model %s", model)
	}

	notchPositions, ok := rotorNotchPositions[model]
	if !ok {
		return Rotor{}, fmt.Errorf("notch position not found for rotor model %s", model)
	}
	rotor := Rotor{mapping, backMapping, position - 'A', offset, notchPositions}
	return rotor, nil
}
