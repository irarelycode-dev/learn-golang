package main

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) setWindowType() {
	i.windowType = "Snow window"
}

func (i *IglooBuilder) setDoorType() {
	i.doorType = "Snow door"
}

func (i *IglooBuilder) setNumFloor() {
	i.floor = 1
}

func (i *IglooBuilder) getHouse() House {
	return House{
		doorType:   i.doorType,
		windowType: i.windowType,
		floor:      i.floor,
	}
}
