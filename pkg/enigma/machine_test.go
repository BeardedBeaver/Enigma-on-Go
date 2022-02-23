package enigma

import (
	"testing"
)

func TestBasicMachine(t *testing.T) {
	machine, err := NewMachine(
		[]string{"I", "II", "III"},
		[]int{'A', 'A', 'A'},
		[]int{0, 0, 0},
		"B",
		[]string{})
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
	machine, err := NewMachine(
		[]string{"I", "II", "III"},
		[]int{'A', 'A', 'A'},
		[]int{0, 0, 0},
		"B",
		[]string{})
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
	machine, err := NewMachine(
		[]string{"II", "V", "VIII"},
		[]int{'S', 'B', 'Z'},
		[]int{4, 3, 2},
		"B",
		[]string{})
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
	machine, err := NewMachine(
		[]string{"I", "V", "VIII"},
		[]int{'D', 'P', 'G'},
		[]int{25, 1, 9},
		"C",
		[]string{"DG", "TH", "PO", "YU"})
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
