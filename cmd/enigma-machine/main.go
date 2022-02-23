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
	machine, err := enigma.NewMachine(
		[]string{rotor1, rotor2, rotor3},
		[]int{int(offset1), int(offset2), int(offset3)},
		[]int{0, 0, 0},
		"B",
		[]string{})
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
