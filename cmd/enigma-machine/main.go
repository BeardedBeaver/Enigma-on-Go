package main

import (
	"EnigmaGo/pkg/enigma"
	"bufio"
	"flag"
	"fmt"
	"os"
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

	fmt.Println("Enter rotor positions (from right to left)")
	fmt.Println("Example: A T H")
	positions, err := readString(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter rotor offsets (from right to left, zero-based)")
	fmt.Println("Example: 6 8 14")
	offsets, err := readString(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter reflector model")
	fmt.Println(enigma.AvailableReflectorModels())
	reflector, err := readString(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	machine, err := enigma.NewMachineFromTextConfig(
		rotors,
		positions,
		offsets,
		reflector,
		[]string{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(machine.PassString(message))
}

func encodeFromArgs() {
	var rotorModels string
	var rotorPositions string
	var rotorOffsets string
	var reflector string
	var message string
	flag.StringVar(&rotorModels, "rotors", "", "list of rotor models right to left")
	flag.StringVar(&rotorPositions, "positions", "A A A", "list of rotor positions right to left")
	flag.StringVar(&rotorOffsets, "offsets", "0 0 0", "list of rotor offsets right to left 0-based")
	flag.StringVar(&reflector, "reflector", "B", "reflector model")
	flag.StringVar(&message, "message", "", "message to encode or decode")
	flag.Parse()

	if len(message) == 0 {
		fmt.Println("Please provide a message to encrypt")
		return
	}
	message = enigma.PreprocessText(message)
	machine, err := enigma.NewMachineFromTextConfig(rotorModels, rotorPositions, rotorOffsets, reflector, []string{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(machine.PassString(message))
}

func main() {
	if len(os.Args) == 1 {
		encodeFromStdin()
	} else {
		encodeFromArgs()
	}
}
