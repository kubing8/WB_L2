package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type logger interface {
	print(message)
}

type message struct {
	level   uint
	payload string
}

type infoLogger struct{ next logger }
type warningLogger struct{ next logger }
type errorLogger struct{ next logger }
type panicLogger struct{}

func (l *infoLogger) print(msg message) {
	if msg.level == 0 {
		fmt.Println("INFO: ", msg.payload)
		return
	}
	l.next.print(msg)
}

func (l *warningLogger) print(msg message) {
	if msg.level == 1 {
		fmt.Println("WARN: ", msg.payload)
		return
	}
	l.next.print(msg)
}

func (l *errorLogger) print(msg message) {
	if msg.level == 2 {
		fmt.Println("ERR: ", msg.payload)
		return
	}
	l.next.print(msg)
}

func (l *panicLogger) print(msg message) {
	fmt.Println("PANIC: ", msg.payload)
}

func main() {
	msg := message{3, "log message"}
	panic := panicLogger{}
	err := errorLogger{&panic}

	warn := warningLogger{&err}
	info := infoLogger{&warn}
	info.print(msg)
}
