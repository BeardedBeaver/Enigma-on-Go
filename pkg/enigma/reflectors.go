package enigma

import (
	"errors"
	"fmt"
)

var reflectorMappings = map[string]string{
	"A":      "EJMZALYXVBWFCRQUONTSPIKHGD",
	"B":      "YRUHQSLDPXNGOKMIEBFZCWVJAT",
	"C":      "FVPJIAOYEDRZXWGCTKUQSBNMHL",
	"B thin": "ENKQAUYWJICOPBLMDXZVFTHRGS",
	"C thin": "RDOBJNTKVEHMLFCWZAXGYIPSUQ",
}

type Reflector struct {
	mapping string
}

func (reflector *Reflector) Pass(character byte) byte {
	code := character - 'A'
	if code < 0 || int(code) >= len(reflector.mapping) {
		code = 0
	}
	return reflector.mapping[code]
}

func NewReflector(model string) (Reflector, error) {
	mapping, ok := reflectorMappings[model]
	if ok == false {
		return Reflector{}, errors.New(fmt.Sprintf("Mapping not found for reflector model %s", model))
	}
	reflector := Reflector{mapping}
	return reflector, nil
}
