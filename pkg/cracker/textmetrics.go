package cracker

// IOC computes Index Of Coincidence metric of a passed
// English text
func IOC(text string) float64 {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var letterCount = make(map[byte]int)
	for _, char := range text {
		letterCount[byte(char)] += 1
	}
	var ioc float64
	textLength := float64(len(text))
	for _, char := range alphabet {
		ioc += float64(letterCount[byte(char)]) / textLength *
			(float64(letterCount[byte(char)]) - 1.0) / (textLength - 1.0)
	}
	return ioc
}
