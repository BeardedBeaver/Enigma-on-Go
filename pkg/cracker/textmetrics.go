package cracker

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// IOC computes Index Of Coincidence metric of a passed text.
// This function will panic if `text` contains anything other
// than an A-Z range
func IOC(text string) float64 {
	var letterCount [26]int
	for _, char := range text {
		letterCount[char-'A'] += 1
	}
	var ioc float64
	textLength := float64(len(text))
	for _, char := range alphabet {
		ioc += float64(letterCount[char-'A']) / textLength *
			(float64(letterCount[char-'A']) - 1.0) / (textLength - 1.0)
	}
	return ioc
}
