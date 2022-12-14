package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Выполнение команд пульта для телевизора

type ICommand interface {
	Execute()
}

type Pult struct {
	command ICommand
}

func (pult *Pult) SetCommand(com ICommand) {
	pult.command = com
}

func (pult *Pult) PressButton() {
	pult.command.Execute()
}

type TV struct {
}

func (tv *TV) SwitchOn() {
	fmt.Println("ТВ включен")
}

func (tv *TV) SwitchOff() {
	fmt.Println("ТВ выключен")
}

type TvOnCommand struct {
	tv TV
}

func (onCom *TvOnCommand) Execute() {
	onCom.tv.SwitchOn()
}

type TvOffCommand struct {
	tv TV
}

func (offCom *TvOffCommand) Execute() {
	offCom.tv.SwitchOn()
}

func main() {
	tv := TV{}
	pult := Pult{}

	pult.SetCommand(&TvOnCommand{tv})
	pult.PressButton()

	pult.SetCommand(&TvOffCommand{tv})
	pult.PressButton()
}
