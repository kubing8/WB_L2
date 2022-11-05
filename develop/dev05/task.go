package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Flags - флаги, поддерживаемые уилитой grep
type Flags struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	Line       bool
}

func main() {
	var (
		afterF      = flag.Int("A", 0, "печатать +N строк после совпадения")
		beforeF     = flag.Int("B", 0, "печатать +N строк до совпадения")
		contextF    = flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
		countF      = flag.Bool("c", false, "количество строк")
		ignoreCaseF = flag.Bool("i", false, "игнорировать регистр")
		invertF     = flag.Bool("v", false, "вместо совпадения, исключать")
		fixedF      = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
		lineF       = flag.Bool("n", false, "напечатать номер строки")
	)
	flag.Parse()
	fl := Flags{After: *afterF, Before: *beforeF, Context: *contextF, Count: *countF,
		IgnoreCase: *ignoreCaseF, Invert: *invertF, Fixed: *fixedF, Line: *lineF}

	var fileName string
	if flag.Arg(0) != "." {
		fileName = flag.Arg(0)
	} else {
		fileName = "C:\\GoCode\\src\\WB_L2\\develop\\dev05\\startText.txt"
	}

	// Открытие файла и загрузка его в память
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

	// Обработка файла
	grepFile(&text, fl, flag.Arg(1))
}

func grepFile(data *[]string, flags Flags, strPattern string) {
	// Проверяем на регистр, если нас он не интересует, то приводим все к меньшему регистру
	if flags.IgnoreCase {
		strPattern = strings.ToLower(strPattern)
	}

	// resIndex заполняется индексами, попадающие под введеные параметры
	resIndex := searchResultIndex(data, flags, strPattern)

	if flags.Line {
		fmt.Print("\nНомера строк: ")
		for _, val := range resIndex {
			fmt.Print(val+1, " ")
		}
	}

	if flags.Count {
		fmt.Print("\nКол-во строк: ", len(resIndex))
	}

	strs := []string{}
	if flags.After > 0 {
		strs = workWithAfter(data, resIndex, flags.After)
	}
	if flags.Before > 0 {
		strs = workWithBefore(data, resIndex, flags.Before)
	}
	if flags.Context > 0 {
		strs = workWithContext(data, resIndex, flags.Context)
	}

	if len(strs) > 0 {
		fmt.Print("\nНайденные интересующие строки: \n")
		for _, str := range strs {
			fmt.Println(str)
		}
	}
}

func workWithAfter(data *[]string, resIndex []int, valFlag int) []string {
	strs := make([]string, 0)
	for _, val := range resIndex {
		strs = append(strs, strsContext(data, val, val+valFlag)...)
	}
	return strs
}
func workWithBefore(data *[]string, resIndex []int, valFlag int) []string {
	strs := make([]string, 0)
	for _, val := range resIndex {
		strs = append(strs, strsContext(data, val-valFlag, val)...)
	}
	return strs
}
func workWithContext(data *[]string, resIndex []int, valFlag int) []string {
	strs := make([]string, 0)
	for _, val := range resIndex {
		strs = append(strs, strsContext(data, val-valFlag, val+valFlag)...)
	}
	return strs
}

func searchResultIndex(data *[]string, flags Flags, strPattern string) []int {
	resIndex := make([]int, 0)

	for ind, row := range *data {
		if flags.IgnoreCase {
			row = strings.ToLower(row)
		}

		// Если нужны те, которые не совпадают - в первый блок, иначе - в else
		if flags.Invert {
			// проверка на точность строки поиска
			if flags.Fixed {
				if row != strPattern {
					resIndex = append(resIndex, ind)
				}

			} else {
				if !strings.Contains(row, strPattern) {
					resIndex = append(resIndex, ind)
				}
			}

		} else {
			// проверка на точность строки поиска
			if flags.Fixed {
				if row == strPattern {
					resIndex = append(resIndex, ind)
				}

			} else {
				if strings.Contains(row, strPattern) {
					resIndex = append(resIndex, ind)
				}
			}
		}
	}
	return resIndex
}

func strsContext(data *[]string, indStart, indEnd int) []string {
	indEnd++ // т.к. край правый не входит
	if indStart < 0 {
		indStart = 0
	}
	if indEnd > len(*data) {
		indEnd = len(*data)
	}
	return (*data)[indStart:indEnd]
}
