package main

import "testing"

type testStruct struct {
	flags   Flags
	outData []string
}

func TestCut(t *testing.T) {
	inputText := readFile("C:\\GoCode\\src\\WB_L2\\develop\\dev06\\startText.txt")
	tests := []testStruct{}

	tempFl := Flags{"1,2", "", false}
	tempOD := []string{
		"drwxr-xr-x      6",
		"drwx------      3",
		"drwxr-xr-x      12",
		"drwx------      5",
		"drwxr-xr-x      7",
		"drwx------      5",
		"drwxr-xr-x      6",
		"drwxr-xr-x      7",
		"drwxr-xr-x      7",
		"drwxr-xr-x      7",
		"drwxr-xr-x      8",
		"drwx------      3",
	}
	tests = append(tests, testStruct{flags: tempFl, outData: tempOD})

	tempFl = Flags{"1", "окт", true}
	tempOD = []string{
		"drwxr-xr-x      7       user    user 4096",
	}
	tests = append(tests, testStruct{flags: tempFl, outData: tempOD})

	for numTest, test := range tests {
		waitData := cutFile(inputText, test.flags)
		for ind, str := range test.outData {
			if str != test.outData[ind] {
				t.Error("Error in ", numTest)
				t.Error(str, " -- ", waitData[ind])
			}
		}
	}
}
