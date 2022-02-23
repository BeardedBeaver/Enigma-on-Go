# Enigma-on-Go

This is an emulation of the Enigma cipher machine written in Go programming language just for fun and to learn Go. 

I'm also going to add a cracker algorithm explained on a Computerphile YouTube channel. 

Created by Dmitriy Linev

Here's a possible list of things that I'm going to do in this project:

 - [x] Fully functional pawl-driven Enigma implementation
   - [x] Rotor model with all main rotors implemented
   - [x] Reflector model with all main reflectors implemented
   - [x] Plugboard
 - [ ] Command-line app for encrypting and decrypting messages
   - [ ] Configuration from arguments
   - [ ] Configuration from standard input
   - [x] Encryption and decryption of a plain text from standard input
   - [ ] Encryption and decryption of a text from a text file
 - [ ] Command-line app for cracking an enigma-encrypted message with an unknown machine config
   - [ ] Bruteforce rotor configuration
   - [x] Index Of Coincidence metric to evaluate a bruteforce result
   - [ ] Mess around with channels and goroutines to speed up a bruteforce
 - [ ] GitHub Actions for testing

---

## Technical details

Rotors I to VIII (Enigma I, M3 and M4) are implemented with wiring according to this wikipedia article: https://en.wikipedia.org/wiki/Enigma_rotor_details#Rotor_wiring_tables

Reflectors A, B, C, B thin and C thin are implemented according to the same wikipedia article. 

---

## Build and run

To build an emulator and a cracker (not implemented yet) run

`go build -o ./bin ./cmd/enigma-machine` or 

`go build -o ./bin ./cmd/enigma-cracker` 

respectively from this project root directory.

To run unittests run `go test ./pkg/enigma ./pkg/cracker` from this project root directory.

GoLand setup I came up with is:

* Build `enigma-machine`
  * Kind: directory
  * Directory: `./cmd/enigma-machine`
  * Output directory: `./bin`
  * Run after build: yes
* Build `enigma-cracker` (Essentially everything is the same except)
  * Directory: `./cmd/enigma-cracker`
* Test
  * Kind: directory
  * Directory: this project root

---

## Other stuff

I recommend this video from Jared Owen https://www.youtube.com/watch?v=ybkkiGtJmkM as it has an extremely good explanation how the machine works. 

Special thanks to Daniel Palloks for explaining the double-stepping effect of the pawl-driven Enigmas. 

---

## License

This project is distributed under MIT license