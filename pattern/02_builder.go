package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
теперь нам требуется построить несколько офисов, в целом они похожи по структуре, но используют разные компоненты

*/

type Office struct {
	windows  string
	door     string
	signaliz string
}

func (o *Office) OfficeToString() string {
	return fmt.Sprintf("Windows: %s,\nDoor: %s,\nSignalization: %s", o.windows, o.door, o.signaliz)
}

type IBuilder interface {
	setWindows()
	setDoor()
	setSignaliz()
	getOffice() Office
}

type Director struct {
	Builder IBuilder
}

func (d *Director) BuildOffice() {
	d.Builder.setWindows()
	d.Builder.setDoor()
	d.Builder.setSignaliz()

	d.Builder.getOffice()
}

type SmallOffice struct {
	windows  string
	door     string
	signaliz string
}

func NewSmallOffice() *SmallOffice {
	return &SmallOffice{}
}

func (so *SmallOffice) setWindows() {
	so.windows = "SmallWindow"
}
func (so *SmallOffice) setDoor() {
	so.door = "SmallDoor"
}
func (so *SmallOffice) setSignaliz() {
	so.signaliz = "SmallSignaliz"
}

func (so *SmallOffice) getOffice() Office {
	return Office{so.windows, so.door, so.signaliz}
}

type LargeOffice struct {
	windows  string
	door     string
	signaliz string
}

func NewLargeOffice() *LargeOffice {
	return &LargeOffice{}
}

func (lo *LargeOffice) setWindows() {
	lo.windows = "largeWindow"
}
func (lo *LargeOffice) setDoor() {
	lo.door = "LargeDoor"
}
func (lo *LargeOffice) setSignaliz() {
	lo.signaliz = "LargeSignaliz"
}

func (lo *LargeOffice) getOffice() Office {
	return Office{lo.windows, lo.door, lo.signaliz}
}

func main() {
	buildSmallOffice := NewSmallOffice()
	directorOfBuilding := Director{buildSmallOffice}
	directorOfBuilding.BuildOffice()
	fmt.Println(buildSmallOffice.getOffice())

	buildLargeOffice := NewLargeOffice()
	directorOfBuilding = Director{buildLargeOffice}
	directorOfBuilding.BuildOffice()
	fmt.Println(buildLargeOffice)

}
