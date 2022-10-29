package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	instr := "eh2k2l\\4n\\\\2dg2"
	outstr, err := decompos(instr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Входная строка: %q\nВыхоная строка: %q\n", instr, outstr)
}

func decompos(str string) (string, error) {
	answer := strings.Builder{}
	for ind, sym := range str {
		//fmt.Println("LOG: ", string(sym))
		if unicode.IsDigit(sym) {
			if ind == 0 {
				return "", errors.New("Incorrect string")
			}

			ch := str[ind-1]
			if unicode.IsDigit(rune(ch)) {
				continue
			}
			if ch == '\\' {
				if str[ind-2] != '\\' {
					num := searchNum(str, ind+1)
					if num == 0 {
						answer.WriteString(string(str[ind]))
					} else {
						for i := 0; i < num; i++ {
							answer.WriteString(string(str[ind]))
						}
					}
					continue
				}
			}

			num := searchNum(str, ind)
			for i := 0; i < num; i++ {
				answer.WriteString(string(ch))
			}
		} else if str[ind] == '\\' {
			continue
		} else if ind == len(str)-1 || !unicode.IsDigit(rune(str[ind+1])) {
			answer.WriteString(string(str[ind]))
		}
	}

	return answer.String(), nil
}

func searchNum(str string, startInd int) int {
	// Узнаем какое это число
	count := 0
	//fmt.Print("LOG: ", str[ind:])
	for _, n := range str[startInd:] {
		//fmt.Println(" num: ", string(n))
		if unicode.IsDigit(n) {
			count++
		} else {
			break
		}
	}

	//fmt.Println(" count: ", count)
	num, _ := strconv.Atoi(str[startInd : startInd+count])
	//fmt.Println(" num: ", num)
	return num
}
