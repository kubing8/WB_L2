package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Рассматриваем опять дома

type IVisitor interface {
	visitHouse(vis *HouseVis)
	visitOffice(vis *OfficeVis)
}

type IBuilding interface {
	accept(v IVisitor)
}

type OfficeVis struct {
	name string
}

func (of *OfficeVis) accept(v IVisitor) {
	v.visitOffice(of)
}

type HouseVis struct {
	name string
}

func (ho *HouseVis) accept(v IVisitor) {
	v.visitHouse(ho)
}

type DirectorVis struct {
	name string
}

func (v *DirectorVis) visitHouse(vis *HouseVis) {
	fmt.Println("Посещаем дом")
}
func (v *DirectorVis) visitOffice(vis *OfficeVis) {
	fmt.Println("Посещаем офис")
}

func main() {
	newHouse := &HouseVis{"NewHouse"}
	newOffice := &OfficeVis{"NewOffice"}

	newDirector := &DirectorVis{"Director"}

	newHouse.accept(newDirector)
	newOffice.accept(newDirector)
}
