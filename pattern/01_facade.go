package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Предположим:
У нас есть офис, и у него есть такие функции: закрыть, створки для окон, включить сигнализацию, закрыть дверь.
И делать это надо в правильном порядке

Для нас все это сложно, мы хотим просто сказать: закрыть. И на этом уйти. Реализуем это через Фасад
*/

type office struct {
	windows  bool
	door     bool
	signaliz bool
}

// Закрытие офиса
func (off *office) CloseWindows() {
	off.windows = false
}

func (off *office) CloseDoor() {
	off.door = false
}

func (off *office) CloseSignaliz() {
	off.signaliz = false
}

// Открытие офиса
func (off *office) OpenWindows() {
	off.windows = true
}

func (off *office) OpenDoor() {
	off.door = true
}

func (off *office) OpenSignaliz() {
	off.signaliz = true
}

// Создадим две функции (фасады), которые замещают нам сложную систему на вызов одной функции
func FacadeCloseOffice(off *office) {
	off.CloseWindows()
	off.CloseSignaliz()
	off.CloseDoor()
}

func FacadeOpenOffice(off *office) {
	off.OpenDoor()
	off.OpenSignaliz()
	off.OpenWindows()
}
