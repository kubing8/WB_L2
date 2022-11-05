package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

/*
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
*/

type testStruct struct {
	f       Flags
	inData  []string
	pattern string
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

func TestGrepAfter(t *testing.T) {
	tests := make([]testStruct, 0)

	inpData := readFile("C:\\GoCode\\src\\WB_L2\\develop\\dev05\\startText.txt")
	inpPattern := "дек"

	tempFlags := Flags{After: 2}
	tempOut := []string{
		"drwxr-xr-x 6 user user 4096 дек 6 14:29 Android",
		"drwx------ 3 user user 4096 янв 14 22:18 Desktop",
		"drwxr-xr-x 12 user user 4096 янв 14 21:49 Documents",
		"drwxr-xr-x 6 user user 4096 дек 6 14:30 Android",
		"drwxr-xr-x 7 user user 4096 июн 10 2015 Sources",
		"drwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBox",
		"дек",
		"январь",
		"фев Дек фен",
	}
	tempTestStruct := testStruct{f: tempFlags, inData: inpData, outData: tempOut, pattern: inpPattern}
	tests = append(tests, tempTestStruct)

	for numTest, test := range tests {
		resIndex := searchResultIndex(&test.inData, test.f, test.pattern)
		test.inData = workWithAfter(&test.inData, resIndex, test.f.After)
		for ind, str := range test.inData {
			if str != test.outData[ind] {
				t.Error("Error in ", numTest)
				t.Error(str, " -- ", test.outData[ind])
			}
		}
	}
}

func TestGrepBefore(t *testing.T) {
	tests := make([]testStruct, 0)

	inpData := readFile("C:\\GoCode\\src\\WB_L2\\develop\\dev05\\startText.txt")
	inpPattern := "дек"

	tempFlags := Flags{Before: 1}
	tempOut := []string{
		"drwxr-xr-x 6 user user 4096 дек 6 14:29 Android",
		"drwx------ 5 user user 12288 янв 15 14:59 Downloads",
		"drwxr-xr-x 6 user user 4096 дек 6 14:30 Android",
		"drwx------ 3 user user 4096 янв 14 22:18 Desktop",
		"дек",
	}
	tempTestStruct := testStruct{f: tempFlags, inData: inpData, outData: tempOut, pattern: inpPattern}
	tests = append(tests, tempTestStruct)

	for numTest, test := range tests {
		resIndex := searchResultIndex(&test.inData, test.f, test.pattern)
		test.inData = workWithBefore(&test.inData, resIndex, test.f.Before)
		for ind, str := range test.inData {
			if str != test.outData[ind] {
				t.Error("Error in ", numTest)
				t.Error(str, " -- ", test.outData[ind])
			}
		}
	}
}

func TestGrepContext(t *testing.T) {
	tests := make([]testStruct, 0)

	inpData := readFile("C:\\GoCode\\src\\WB_L2\\develop\\dev05\\startText.txt")
	inpPattern := "дек"

	tempFlags := Flags{Context: 1}
	tempOut := []string{
		"drwxr-xr-x 6 user user 4096 дек 6 14:29 Android",
		"drwx------ 3 user user 4096 янв 14 22:18 Desktop",
		"drwx------ 5 user user 12288 янв 15 14:59 Downloads",
		"drwxr-xr-x 6 user user 4096 дек 6 14:30 Android",
		"drwxr-xr-x 7 user user 4096 июн 10 2015 Sources",
		"drwx------ 3 user user 4096 янв 14 22:18 Desktop",
		"дек",
		"январь",
	}

	tempTestStruct := testStruct{f: tempFlags, inData: inpData, outData: tempOut, pattern: inpPattern}
	tests = append(tests, tempTestStruct)

	for numTest, test := range tests {
		resIndex := searchResultIndex(&test.inData, test.f, test.pattern)
		test.inData = workWithContext(&test.inData, resIndex, test.f.Context)
		for ind, str := range test.inData {
			if str != test.outData[ind] {
				t.Error("Error in ", numTest)
				t.Error(str, " -- ", test.outData[ind])
			}
		}
	}
}
