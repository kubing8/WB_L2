package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	fmt.Println("Hello")
	t, err := ntp.Time("ru.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Current time (time.Now): %v\nCurrent time (NTP): %v",
		TimeFormat(time.Now()), TimeFormat(t))
}

func TimeFormat(time time.Time) string {
	return time.Format("15:04:05")
}
