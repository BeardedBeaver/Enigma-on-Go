package enigma

import (
	"encoding/json"
	"strings"
	"unicode"
)

// PreprocessText converts passed text to upper case
// and removes all except actual letters to prepare
// it to be passed to the Enigma
func PreprocessText(message string) string {
	result := strings.ToUpper(message)

	// remove unwanted characters
	var b strings.Builder
	b.Grow(len(result))
	for _, ch := range result {
		if unicode.IsLetter(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

type RotorConfig struct {
	Model    string `json:"model"`
	Position int    `json:"position"`
	Offset   int    `json:"offset"`
}

type MachineConfig struct {
	RotorConfig       []RotorConfig `json:"rotorConfig"`
	ReflectorModel    string        `json:"reflectorModel"`
	PlugboardMappings []string      `json:"plugboardMappings,omitempty"`
}

type Machine struct {
	refl    Reflector
	rot     []Rotor
	plugbrd Plugboard
	conf    MachineConfig
}

func (machine *Machine) passChar(character byte) byte {

	// first we check which rotors to spin
	rotorsToSpin := []int{0}
	for i := range machine.rot {
		if i != len(machine.rot)-1 {
			if machine.rot[i].IsAtNotch() {
				rotorsToSpin = append(rotorsToSpin, i+1)

				// here the double-stepping is performed - the effect on pawl-driven Enigmas
				// that causes the rotor to the right also spin with the current one
				// because of the design of the notches;
				// first we check if rotor to the right (with smaller index) isn't supposed to spin already
				// and if it's not, we add it to the slice
				alreadyHasToSpin := false
				for _, index := range rotorsToSpin {
					if index == i {
						alreadyHasToSpin = true
						break
					}
				}
				if alreadyHasToSpin == false {
					rotorsToSpin = append(rotorsToSpin, i)
				}
			}
		}
	}

	// and now spin
	for i := range rotorsToSpin {
		machine.rot[i].Spin()
	}

	// forward plugboard pass
	character = machine.plugbrd.PassCharacter(character)

	// forward rotors pass
	for _, rotor := range machine.rot {
		character = rotor.PassForward(character)
	}

	// reflector
	character = machine.refl.Pass(character)

	// backward rotors pass
	for i := len(machine.rot) - 1; i >= 0; i-- {
		character = machine.rot[i].PassBackward(character)
	}

	// backward plugboard pass
	character = machine.plugbrd.PassCharacter(character)

	return character
}

func (machine *Machine) PassString(message string) string {
	var result string
	for _, character := range message {
		encodedCharacter := machine.passChar(byte(character))
		result += string(encodedCharacter)
	}
	return result
}

// NewMachine creates a new Enigma machine.
// Rotor models are listed from right to left
func NewMachine(conf MachineConfig) (Machine, error) {
	var machine Machine
	machine.conf = conf
	for _, config := range conf.RotorConfig {
		rotor, err := NewRotor(config.Model, config.Position, config.Offset)
		if err != nil {
			return Machine{}, err
		}
		machine.rot = append(machine.rot, rotor)
	}
	refl, err := NewReflector(conf.ReflectorModel)
	if err != nil {
		return Machine{}, err
	}
	machine.refl = refl

	plugbrd, err := NewPlugboard(conf.PlugboardMappings)
	if err != nil {
		return Machine{}, err
	}
	machine.plugbrd = plugbrd
	return machine, nil
}

func NewMachineFromJson(j []byte) (Machine, error) {
	conf := MachineConfig{}
	err := json.Unmarshal(j, &conf)
	if err != nil {
		return Machine{}, err
	}
	return NewMachine(conf)
}
