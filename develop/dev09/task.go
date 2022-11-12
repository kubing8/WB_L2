package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Flags - флаги, поддерживаемые уилитой wget
type Flags struct {
	OutDir string
}

// Наверное нужна опци -О => директория вывода файла, если не указана, то выводить в текущую
func main() {
	outDir := flag.String("O", ".", "Use to specify the output directory")
	flag.Parse()

	fl := Flags{*outDir}
	url := flag.Arg(0)

	page := wget(url)
	writeFile(fl, url, page)
	//fmt.Print(string(page))
}

func writeFile(fl Flags, url string, strs []byte) {
	var path string
	url = strings.TrimPrefix(url, "https://")
	domens := strings.Split(url, "/")
	file := domens[len(domens)-1]

	if fl.OutDir == "." {
		path = "C:\\GoCode\\src\\WB_L2\\develop\\dev09\\" + file + ".html"
	} else {
		path = fl.OutDir + file + ".html"
	}

	if err := os.WriteFile(path,
		strs, os.ModePerm); err != nil {
		fmt.Println("Ошибка записи в файл.\n", err)
	}
}

func wget(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body
}
