package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

type testStruct struct {
	f       Flags
	inData  []string
	outData []string
}

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error open file!\n", err)
	}
	defer file.Close()

	var text []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}

func TestSort(t *testing.T) {
	tests := make([]testStruct, 3)
	tr, fa := true, false

	tests[0].inData = readFile("C:\\GoCode\\src\\WB_L2\\develop\\dev03\\startText.txt")
	tests[0].f = Flags{1, &fa, &tr, &fa}
	tests[0].outData = []string{
		"drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks",
		"drwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBox",
		"drwxr-xr-x 7 user user 4096 июн 10 2015 Sources",
		"drwxr-xr-x 6 user user 4096 дек 6 14:30 Android",
		"drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks",
		"drwxr-xr-x 12 user user 4096 янв 14 21:49 Documents",
		"drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures",
		"drwxr-xr-x 6 user user 4096 дек 6 14:29 Android",
		"drwx------ 5 user user 12288 янв 15 14:59 Downloads",
		"drwx------ 5 user user 12288 янв 15 14:59 Downloads",
		"drwx------ 3 user user 4096 янв 14 22:18 Desktop",
		"drwx------ 3 user user 4096 янв 14 22:18 Desktop",
	}

	tests[1].inData = readFile("C:\\GoCode\\src\\WB_L2\\develop\\dev03\\startText.txt")
	tests[1].f = Flags{2, &tr, &fa, &fa}
	tests[1].outData = []string{
		"drwx------ 3 user user 4096 янв 14 22:18 Desktop",
		"drwx------ 3 user user 4096 янв 14 22:18 Desktop",
		"drwx------ 5 user user 12288 янв 15 14:59 Downloads",
		"drwx------ 5 user user 12288 янв 15 14:59 Downloads",
		"drwxr-xr-x 6 user user 4096 дек 6 14:29 Android",
		"drwxr-xr-x 6 user user 4096 дек 6 14:30 Android",
		"drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks",
		"drwxr-xr-x 7 user user 4096 июн 10 2015 Sources",
		"drwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBox",
		"drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks",
		"drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures",
		"drwxr-xr-x 12 user user 4096 янв 14 21:49 Documents",
	}

	// Проверка на -u удаление дубликатов
	tests[2].inData = readFile("C:\\GoCode\\src\\WB_L2\\develop\\dev03\\startText.txt")
	tests[2].f = Flags{K: -1, N: &fa, R: &fa, U: &tr}
	tests[2].outData = []string{
		"drwxr-xr-x 6 user user 4096 дек 6 14:29 Android",
		"drwx------ 3 user user 4096 янв 14 22:18 Desktop",
		"drwxr-xr-x 12 user user 4096 янв 14 21:49 Documents",
		"drwx------ 5 user user 12288 янв 15 14:59 Downloads",
		"drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks",
		"drwxr-xr-x 6 user user 4096 дек 6 14:30 Android",
		"drwxr-xr-x 7 user user 4096 июн 10 2015 Sources",
		"drwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBox",
	}

	for numTest, test := range tests {
		sortFile(&test.inData, test.f)
		for ind, str := range test.inData {
			if str != test.outData[ind] {
				t.Error("Error in ", numTest)
				t.Error(str, " -- ", test.outData[ind])
			}
		}
	}

}
