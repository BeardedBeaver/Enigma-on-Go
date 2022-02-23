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
