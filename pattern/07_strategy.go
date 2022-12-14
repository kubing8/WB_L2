package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type ITransp interface {
	Go()
}

type Car struct {
}

func (car *Car) Go() {
	fmt.Println("Едем на машине")
}

type Camel struct {
}

func (camel *Camel) Go() {
	fmt.Println("Идем верхом")
}

type Client struct {
	name    string
	strateg ITransp
}

func (c *Client) NewClient(str string, tran ITransp) {
	c.name = str
	c.strateg = tran
}

func (c *Client) ChangeMove(tran ITransp) {
	c.strateg = tran
}

func (c *Client) Move() {
	c.strateg.Go()
}

func main() {
	camel := &Camel{}
	car := &Car{}

	me := Client{}
	me.NewClient("Ivan", car)
	me.Move()

	me.ChangeMove(camel)
	me.Move()
}
