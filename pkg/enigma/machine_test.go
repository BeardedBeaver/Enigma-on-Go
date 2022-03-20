package enigma

import (
	"testing"
)

func TestBasicMachine(t *testing.T) {
	config := []RotorConfig{
		{"I", 'A', 0},
		{"II", 'A', 0},
		{"III", 'A', 0},
	}
	machine, err := NewMachine(config, "B", []string{})
	if err != nil {
		t.Error("Can't create Machine: ", err)
	}
	message := machine.PassString("HELLOWORLD")
	actual := "MFNCZBBFZM"
	if message != actual {
		t.Error("Error encrypting message - expected:\n", actual, "\ngot:\n", message)
	}
}

func TestLongerStringMachine(t *testing.T) {
	config := []RotorConfig{
		{"I", 'A', 0},
		{"II", 'A', 0},
		{"III", 'A', 0},
	}
	machine, err := NewMachine(config, "B", []string{})
	if err != nil {
		t.Error("Can't create Machine: ", err)
	}

	message := machine.PassString("HIIAMASIMPLEENIGMAEMULATORBUILTINGOJUSTFORFUN" +
		"ANDTHISISANEXAMPLEOFAVERYLONGSTRINGTOMAKESURETHATLONGSTRINGS" +
		"CAUSESECONDANDPROBABLYTHIRDROTORTOSPINATTHENOTCHINTHERIGHTTIME")

	actual := "MQJMFIAZTACMSEZPCCWZJURKJWTIQTRPDJQCGEWHHYICLCYUATCBCDVSYP" +
		"VKTCABARWFLLKAIETENAQHPQBDVRKXMUVTHVFKSXZLBLIIUKBRHLMAQWJHWQMDDG" +
		"KBEMSNWNXCKXJXZMGOKGZVRBSTOMGZEOZGFCHSNYQQOLV"
	if message != actual {
		t.Error("Error encrypting message - expected:\n", actual, "\ngot:\n", message)
	}

}

func TestRingOffset(t *testing.T) {
	config := []RotorConfig{
		{"II", 'S', 4},
		{"V", 'B', 3},
		{"VIII", 'Z', 2},
	}
	machine, err := NewMachine(config, "B", []string{})
	if err != nil {
		t.Error("Can't create Machine: ", err)
	}
	output := machine.PassString("MYENIGMAWORKSPERFECTLYWITHRINGSETTINGSMEOW")
	expected := "ETJDZJCGFGQMUGHNAHUJFKBSSLMQDMVSHWUDELPFQT"
	if output != expected {
		t.Error("Error encrypting message - expected:\n", expected, "\ngot:\n", output)
	}
}

func TestMachine_PassString_AllStuff(t *testing.T) {
	config := []RotorConfig{
		{"I", 'D', 25},
		{"V", 'P', 1},
		{"VIII", 'G', 9},
	}
	machine, err := NewMachine(config, "C", []string{"DG", "TH", "PO", "YU"})
	if err != nil {
		t.Error("Can't create Machine: ", err)
	}

	message := "OHWOWMYENIGMAWORKSEVENWITHAPLUGBOARDANDSTUFFHOWCOOLISTHAT"
	output := machine.PassString(message)
	expected := "VNBWAIQVVFZIGEIGYPVWWREKYFLSEAAXBYVVSBXMKSXPRYPEGYKQCDAWL"
	if output != expected {
		t.Error("Error encrypting message - expected:\n", expected, "\ngot:\n", output)
	}
}
