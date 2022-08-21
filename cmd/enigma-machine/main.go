package main

import (
	"EnigmaGo/pkg/enigma"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readString(reader *bufio.Reader) (string, error) {
	result, err := reader.ReadString('\n')
	result = strings.Trim(result, "\n")
	if err != nil {
		return "", err
	}
	return result, nil
}

// encodeFromStdin is used for both encoding and decoding a text with a known
// Enigma configuration
func encodeFromStdin() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a message:")
	message, err := readString(in)
	if err != nil {
		fmt.Println(err)
		return
	}
	message = enigma.PreprocessText(message)

	fmt.Println("Enter rotor models (from right to left)")
	fmt.Println("Example: I II III")
	fmt.Println(enigma.AvailableRotorModels())
	rotors, err := readString(in)
	if err != nil {
		fmt.Println(err)
		return
	}
	rotorModels := strings.Split(rotors, " ")

	fmt.Println("Enter rotor positions (from right to left)")
	fmt.Println("Example: A T H")
	positions, err := readString(in)
	if err != nil {
		fmt.Println(err)
		return
	}
	rotorPositions := strings.Split(positions, " ")
	if len(rotorPositions) != len(rotorModels) {
		fmt.Println("Number of rotor models should be the same as the number of rotor positions")
		return
	}

	fmt.Println("Enter rotor offsets (from right to left, zero-based)")
	fmt.Println("Example: 6 8 14")
	offsets, err := readString(in)
	if err != nil {
		fmt.Println(err)
		return
	}
	rotorOffsets := strings.Split(offsets, " ")
	if len(rotorOffsets) != len(rotorModels) {
		fmt.Println("Number of rotor offsets should be the same as the number of rotor models")
		return
	}

	rotorConfig := make([]enigma.RotorConfig, 0, len(rotorModels))
	for i := 0; i < len(rotorModels); i++ {
		rotorSetting := int(rotorPositions[i][0])
		rotorOffset, err := strconv.Atoi(rotorOffsets[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		rotorConfig = append(
			rotorConfig,
			enigma.RotorConfig{Model: rotorModels[i], Position: rotorSetting, Offset: rotorOffset},
		)
	}

	fmt.Println("Enter reflector model")
	fmt.Println(enigma.AvailableReflectorModels())
	reflectorModel, err := readString(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	encoded, err := encode(message, rotorConfig, reflectorModel, []string{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encoded)
}

func encode(message string,
	rotorConfig []enigma.RotorConfig,
	reflectorModel string,
	plugboardMapping []string) (string, error) {
	machineConfig := enigma.MachineConfig{
		RotorConfig:       rotorConfig,
		ReflectorModel:    reflectorModel,
		PlugboardMappings: plugboardMapping,
	}
	machine, err := enigma.NewMachine(machineConfig)
	if err != nil {
		return "", err
	}
	return machine.PassString(message), nil
}

func main() {
	encodeFromStdin()
}
