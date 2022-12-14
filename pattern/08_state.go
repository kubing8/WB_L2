package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/State_pattern
*/
type Status interface {
	connect()
}
type Man struct {
	sleepStatus Status
	workStatus  Status
	walkStatus  Status

	currentStatus Status
}

func (m *Man) connect() {
	switch m.currentStatus {
	case m.sleepStatus:
		m.currentStatus.connect()
		m.setState(m.workStatus)
	case m.workStatus:
		m.currentStatus.connect()
		m.setState(m.walkStatus)
	case m.walkStatus:
		m.currentStatus.connect()
		m.setState(m.sleepStatus)
	}

}
func (m *Man) setState(s Status) {
	m.currentStatus = s
}
func newMan() *Man {
	return &Man{sleepStatus: &Sleep{}, workStatus: &Work{}, walkStatus: &Walk{}}
}

type Sleep struct{}

func (s *Sleep) connect() {
	fmt.Println("Я сплю. Позвоните позднее!")
}

type Work struct{}

func (w *Work) connect() {
	fmt.Println("Я на работе! Свяжитесь со мной через email!")
}

type Walk struct{}

func (w *Walk) connect() {
	fmt.Println("Я на связи!")
}
func main() {
	man := newMan()
	man.setState(&Work{})
	for i := 0; i < 7; i++ {
		man.connect()
	}
}
