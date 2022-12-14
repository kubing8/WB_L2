package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type ITransport interface {
	SetName(string)
	GetName() string
}

type Transort struct {
	name string
}

func (t *Transort) SetName(str string) {
	t.name = str
}

func (t *Transort) GetName() string {
	return t.name
}

// Для скутера
type scuter struct {
	Transort
}

func NewScuter() ITransport {
	return &scuter{Transort{name: "Scuter 2.0"}}
}

// Используем такси
type taxi struct {
	Transort
}

func NewTaxi() ITransport {
	return &taxi{Transort{name: "Almera A307BB"}}
}

func GetTransort(str string) ITransport {
	if str == "Taxi" {
		return NewTaxi()
	} else if str == "Scuter" {
		return NewScuter()
	}
	return nil
}

func main() {
	scut := GetTransort("Scuter")
	tax := GetTransort("Taxi")

	fmt.Println(scut, " ", tax)
}
