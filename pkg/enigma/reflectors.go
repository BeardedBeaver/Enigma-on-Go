package enigma

import (
	"fmt"
)

var reflectorMappings = map[string]string{
	"A":      "EJMZALYXVBWFCRQUONTSPIKHGD",
	"B":      "YRUHQSLDPXNGOKMIEBFZCWVJAT",
	"C":      "FVPJIAOYEDRZXWGCTKUQSBNMHL",
	"B thin": "ENKQAUYWJICOPBLMDXZVFTHRGS",
	"C thin": "RDOBJNTKVEHMLFCWZAXGYIPSUQ",
}

func AvailableReflectorModels() []string {
	return []string{
		"A",
		"B",
		"C",
		"B thin",
		"C thin",
	}
}

type Reflector struct {
	mapping string
}

func (reflector *Reflector) Pass(character byte) byte {
	code := character - 'A'
	if int(code) >= len(reflector.mapping) {
		code = 0
	}
	return reflector.mapping[code]
}

func NewReflector(model string) (Reflector, error) {
	mapping, ok := reflectorMappings[model]
	if !ok {
		return Reflector{}, fmt.Errorf("mapping not found for reflector model %s", model)
	}
	reflector := Reflector{mapping}
	return reflector, nil
}
