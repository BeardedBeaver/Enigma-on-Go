package main

import (
	"EnigmaGo/pkg/cracker"
	"EnigmaGo/pkg/enigma"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime/pprof"
	"sort"
	"sync"
	"time"
)

// printResult function sorts an input map by value and outputs top scores
func printResult(scores map[string]float64, top int) {
	type KeyValuePair struct {
		Key   string
		Value float64
	}

	var pairs []KeyValuePair
	for k, v := range scores {
		pairs = append(pairs, KeyValuePair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	for index, kv := range pairs {
		if index > top {
			break
		}
		fmt.Printf("%s, %f\n", kv.Key, kv.Value)
	}
}

func decodeRotorConfig(rotorModels []string, message string) (map[string]float64, error) {
	scores := make(map[string]float64)
	if len(rotorModels) != 3 {
		return scores, errors.New("not found")
	}
	for r1 := 0; r1 < 26; r1++ {
		for r2 := 0; r2 < 26; r2++ {
			for r3 := 0; r3 < 26; r3++ {
				rotorConfig := []enigma.RotorConfig{
					{Model: rotorModels[0], Position: 'A' + r1, Offset: 0},
					{Model: rotorModels[1], Position: 'A' + r2, Offset: 0},
					{Model: rotorModels[2], Position: 'A' + r3, Offset: 0},
				}
				machineConfig := enigma.MachineConfig{
					RotorConfig:       rotorConfig,
					ReflectorModel:    "B",
					PlugboardMappings: []string{},
				}

				machine, err := enigma.NewMachine(machineConfig)
				if err != nil {
					return scores, fmt.Errorf("decode rotor config error: %v", err)
				}
				decoded := machine.PassString(message)
				ioc := cracker.IOC(decoded)
				key, err := json.Marshal(machineConfig)
				if err != nil {
					return scores, err
				}
				scores[string(key[:])] = ioc
			}
		}
	}
	return scores, nil
}

// tries to guess the Enigma configuration and decode passed text
func decode(message string) {
	start := time.Now()

	f, err := os.Create("profile.prof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	message = enigma.PreprocessText(message)
	rotorModels := enigma.AvailableRotorModels()

	permutationSender := make(chan []int)
	scoreSender := make(chan map[string]float64)

	go cracker.GeneratePermutations(8, 3, permutationSender)

	wg := sync.WaitGroup{}
	for permutation := range permutationSender {
		rotorConfig := make([]string, 0, 3)
		for _, index := range permutation {
			rotorConfig = append(rotorConfig, rotorModels[index])
		}
		wg.Add(1)
		go func(config []string, message string) {
			defer wg.Done()
			result, err := decodeRotorConfig(config, message)
			if err != nil {
				fmt.Println(err) // panic?
				return
			}
			fmt.Println(config)
			scoreSender <- result
		}(rotorConfig, message)
	}

	go func() {
		wg.Wait()
		close(scoreSender)
	}()

	scores := make(map[string]float64)
	for score := range scoreSender {
		for k, v := range score {
			scores[k] = v
		}
	}

	printResult(scores, 10)
	fmt.Printf("Elapsed time %s\n", time.Since(start))
}

// Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.
// Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.

// 4 5 3 IDC 22 6 13

// fwjin qmteg xvvpw frdat rbaus kgufw yicto cppif lgjqo iuskb cqzoj nwoxk xwgkw mlvwm nijqi bojma liixl bveax zxvmf fdewr hcrvw wruhj zlzfp fqius ojmzv cxrrg hdhyt chuqi tgfcs nieit jpoph kegfj szppy nnenm bjsgc podwv upmyc quigk psbki sjszb dn
func main() {
	fmt.Println("Enter a message:")
	in := bufio.NewReader(os.Stdin)
	message, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	decode(message)
}
