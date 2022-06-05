package main

import (
	"EnigmaGo/pkg/enigma"
	"bufio"
	"fmt"
	"os"
)

// encode is used for both encoding and decoding a text with a known
// Enigma configuration
func encode() {
	fmt.Println("Enter a message:")
	in := bufio.NewReader(os.Stdin)
	message, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	message = enigma.PreprocessText(message)

	fmt.Println("Enter rotor models (from right to left)")
	var rotor1, rotor2, rotor3 string
	_, err = fmt.Scan(&rotor1, &rotor2, &rotor3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter rotor settings (from right to left)")
	var offset1, offset2, offset3 byte
	_, err = fmt.Scanf("%c %c %c", &offset1, &offset2, &offset3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s/%s/%s\n", rotor1, rotor2, rotor3)
	rotorConfig := []enigma.RotorConfig{
		{rotor1, int(offset1), 0},
		{rotor2, int(offset2), 0},
		{rotor3, int(offset3), 0},
	}
	machineConfig := enigma.MachineConfig{
		RotorConfig:       rotorConfig,
		ReflectorModel:    "B",
		PlugboardMappings: []string{},
	}
	machine, err := enigma.NewMachine(machineConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	encoded := machine.PassString(message)
	fmt.Println(encoded)
}

func main() {
	encode()
}
