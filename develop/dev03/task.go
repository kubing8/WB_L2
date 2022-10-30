package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное

Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Flags - флаги, поддерживаемые уилитой sort
type Flags struct {
	K int   // сортировка по колонкам
	N *bool // сортировка по числвой колонке (целые\дробные). Обязательно указание ключа -k
	R *bool // сортировка в обратном порядке
	U *bool // не выводить повторяюиеся строки
}

func main() {
	var (
		collumn int   // сортировка по колонкам
		num     *bool // сортировка по числу
		revSort *bool // сортировка в обратном порядке
		uniqStr *bool // не выводить повторяюиеся строки
	)
	flag.IntVar(&collumn, "k", -1, "Number of column for sorting")
	num = flag.Bool("n", false, "Sorting by numeric column, it is mandatory to specify the key -k")
	revSort = flag.Bool("r", false, "Sorting in revers order, it is mandatory to specify the key -k")
	uniqStr = flag.Bool("u", false, "Delete repeatable rows")

	flag.Parse()

	fl := Flags{K: collumn, N: num, R: revSort, U: uniqStr}

	// Узнаем имя файла, над которым будут производиться действия
	var fileName string
	if flag.Arg(0) != "" {
		fileName = flag.Arg(0)
	} else {
		fileName = "C:\\GoCode\\src\\WB_L2\\develop\\dev03\\startText.txt"
	}

	// Открытие файла и загрузка его в ОП
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
	sortFile(&text, fl)

	// Преобразуем в []byte
	// Возврат результата в файл
	if err := os.WriteFile("C:\\GoCode\\src\\WB_L2\\develop\\dev03\\finalText.txt",
		[]byte(strings.Join(text, "\n")), os.ModePerm); err != nil {
		fmt.Println("Ошибка записи в файл.\n", err)
	}
}

func sortFile(data *[]string, flags Flags) {
	if *flags.U {
		remDup(data)
	}

	if flags.K != -1 {
		flags.K-- // Для указания индекса колонки, а не номера
		if *flags.N {
			sortColumnNum(data, flags.K)
		} else {
			sortColumnString(data, flags.K)
		}
	}

	if *flags.R {
		reverseStringSlice(data)
	}
}

func remDup(data *[]string) {
	setStr := make(map[int]struct{})
	setInd := make(map[int]struct{})

	// Находим все повторяющиеся индексы
	for ind, str := range *data {
		hash := hashSumBytes(str)
		if _, ok := setStr[hash]; ok {
			setInd[ind] = struct{}{}
		} else {
			setStr[hash] = struct{}{}
		}
	}
	delCount := 0
	for k := range setInd {
		k -= delCount // смещение, от уже удаленной строки
		if k == len(*data)-1 {
			*data = (*data)[:k-1]
		} else {
			*data = append((*data)[:k], (*data)[k+1:]...)
		}
		delCount++ // Добавляем к смещению, т.к. только что удалили один элемент
	}
}

// Нахождение хэш-суммы строки
func hashSumBytes(str string) int {
	sum := 0
	for ind, ch := range str {
		sum += ind * int(ch)
	}
	return sum
}

// Сортирует сткроки по текстовым колонкам, если укзаать вместе флагом -n, то отсортирут по числовому значению
func sortColumnString(data *[]string, column int) {
	mCol := make(map[string]int)
	for ind, str := range *data {
		words := strings.Split(str, " ")
		if column >= len(words) {
			fmt.Println("Неверное указание колонки для сортировки! ")
			os.Exit(1)
		}
		words[column] += strconv.Itoa(ind) // Т.к. мы делаем из этого ключ, мы добавляем индекс строки, чтобы все были разыне
		mCol[words[column]] = ind
	}

	keys := make([]string, 0, len(mCol))
	for k := range mCol {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	retData := []string{}
	for _, val := range keys {
		//fmt.Println((*data)[mCol[val]])
		retData = append(retData, (*data)[mCol[val]])
	}

	*data = retData
}

// Сортирует слайс строк по числовому полю (целые\дробные числа)
func sortColumnNum(data *[]string, column int) {
	type structForSort struct {
		col float64
		ind int
	}
	answerStruct := []structForSort{}
	for ind, str := range *data {
		words := strings.Split(str, " ")
		if column >= len(words) {
			fmt.Println("Неверное указание колонки для сортировки! ")
			os.Exit(1)
		}
		word := strings.Split(words[column], ".")
		var num float64
		if len(word) > 1 {
			iPart, _ := strconv.Atoi(word[0])
			zPart, _ := strconv.Atoi(word[1])
			num = float64(iPart) + float64(zPart/len(word[1]))
		} else {
			numInt, _ := strconv.Atoi(word[0])
			num = float64(numInt)
		}
		answerStruct = append(answerStruct, structForSort{num, ind})
	}

	sort.Slice(answerStruct, func(i, j int) (less bool) {
		return answerStruct[i].col < answerStruct[j].col
	})

	retData := []string{}
	for _, val := range answerStruct {
		retData = append(retData, (*data)[val.ind])
	}

	*data = retData
}

// Переворачивает слайс строк
func reverseStringSlice(data *[]string) {
	i := 0
	j := len(*data) - 1
	for i < j {
		(*data)[i], (*data)[j] = (*data)[j], (*data)[i]
		i++
		j--
	}
}
