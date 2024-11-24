package cracker

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestIOC(t *testing.T) {
	text := "To be or not to be that is the question" +
		"Whether tis Nobler in the mind to suffer" +
		"The Slings and Arrows of outrageous Fortune" +
		"Or to take Arms against a Sea of troubles" +
		"And by opposing end them" +
		"William Shakespeare Hamlet"

	text = strings.ToUpper(text)
	text = strings.ReplaceAll(text, " ", "")
	ioc := IOC(text)
	var actualIOC float64 = 0.06773
	if math.Abs(ioc-actualIOC) > 1e-5 {
		message, _ := fmt.Printf("Error computing IOC, expected %f, got %f", actualIOC, ioc)
		t.Error(message)
	}

	text = "BUQERUZAWTMJWSFLAMVDPELPNQJURBECYTNGJOFPFNZUHMNTDCSRYDWIHFTKPVDRCLEEUFAHOIEFGZEKRUMNLTHDYJORHNMXEZMX"
	ioc = IOC(text)
	actualIOC = 0.03636
	if math.Abs(ioc-actualIOC) > 1e-5 {
		message := fmt.Sprintf("Error computing IOC, expected %f, got %f", actualIOC, ioc)
		t.Error(message)
	}
}

func iocOriginal(text string) float64 {
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

func iocNoConversion(text string) float64 {
	var letterCount = make(map[rune]int)
	for _, char := range text {
		letterCount[char] += 1
	}
	var ioc float64
	textLength := float64(len(text))
	for _, char := range alphabet {
		ioc += float64(letterCount[char]) / textLength *
			(float64(letterCount[char]) - 1.0) / (textLength - 1.0)
	}
	return ioc
}

func BenchmarkIocOriginal(b *testing.B) {
	longText := "THISISAMUCHLONGERTEXTDESIGNEDTOSTRESSTHETIOCFUNCTIONITHASENOUGHCONTENTTOEVALUATEPERFORMANCEUNDERLARGERINPUTSIZESTHISSHOULDDIVEAGOODIDEAOFHOWTHEFUNCTIONSCALESWITHLARGERSTRINGSTHETEXTSIZEINCREASES"

	b.ResetTimer() // Reset timer to exclude setup time
	for i := 0; i < b.N; i++ {
		iocOriginal(longText)
	}
}

func BenchmarkIocNoConversion(b *testing.B) {
	longText := "THISISAMUCHLONGERTEXTDESIGNEDTOSTRESSTHETIOCFUNCTIONITHASENOUGHCONTENTTOEVALUATEPERFORMANCEUNDERLARGERINPUTSIZESTHISSHOULDDIVEAGOODIDEAOFHOWTHEFUNCTIONSCALESWITHLARGERSTRINGSTHETEXTSIZEINCREASES"

	b.ResetTimer() // Reset timer to exclude setup time
	for i := 0; i < b.N; i++ {
		iocNoConversion(longText)
	}
}

func BenchmarkIoc(b *testing.B) {
	longText := "THISISAMUCHLONGERTEXTDESIGNEDTOSTRESSTHETIOCFUNCTIONITHASENOUGHCONTENTTOEVALUATEPERFORMANCEUNDERLARGERINPUTSIZESTHISSHOULDDIVEAGOODIDEAOFHOWTHEFUNCTIONSCALESWITHLARGERSTRINGSTHETEXTSIZEINCREASES"

	b.ResetTimer() // Reset timer to exclude setup time
	for i := 0; i < b.N; i++ {
		IOC(longText)
	}
}
