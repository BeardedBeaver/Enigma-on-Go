package enigma

import (
	"testing"
)

func TestBasicMachine(t *testing.T) {
	rotorConfig := []RotorConfig{
		{"I", 'A', 0},
		{"II", 'A', 0},
		{"III", 'A', 0},
	}
	machineConfig := MachineConfig{
		rotorConfig,
		"B",
		[]string{},
	}
	machine, err := NewMachine(machineConfig)
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
	rotorConfig := []RotorConfig{
		{"I", 'A', 0},
		{"II", 'A', 0},
		{"III", 'A', 0},
	}
	machineConfig := MachineConfig{
		rotorConfig,
		"B",
		[]string{},
	}
	machine, err := NewMachine(machineConfig)
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
	rotorConfig := []RotorConfig{
		{"II", 'S', 4},
		{"V", 'B', 3},
		{"VIII", 'Z', 2},
	}
	machineConfig := MachineConfig{
		rotorConfig,
		"B",
		[]string{},
	}
	machine, err := NewMachine(machineConfig)
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
	rotorConfig := []RotorConfig{
		{"I", 'D', 25},
		{"V", 'P', 1},
		{"VIII", 'G', 9},
	}
	machineConfig := MachineConfig{
		rotorConfig,
		"C",
		[]string{"DG", "TH", "PO", "YU"},
	}
	machine, err := NewMachine(machineConfig)
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

func TestNewMachineFromJson(t *testing.T) {
	config := []byte(`
	{
		"rotorConfig": [
		{"model": "I", "position": 68, "offset": 25},
		{"model": "V", "position": 80, "offset": 1},
		{"model": "VIII", "position": 71, "offset": 9}
		],
		"reflectorModel": "C",
		"plugboardMappings": ["DG", "TH", "PO", "YU"]
	}
	`)
	machine, err := NewMachineFromJson(config)
	if err != nil {
		t.Error(err)
	}
	message := "OHWOWMYENIGMAWORKSEVENWITHAPLUGBOARDANDSTUFFHOWCOOLISTHAT"
	output := machine.PassString(message)
	expected := "VNBWAIQVVFZIGEIGYPVWWREKYFLSEAAXBYVVSBXMKSXPRYPEGYKQCDAWL"
	if output != expected {
		t.Error("Error encrypting message - expected:\n", expected, "\ngot:\n", output)
	}
}
