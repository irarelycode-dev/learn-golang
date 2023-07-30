package main

type IPlayer interface {
	setName(name string)
	setAge(age int)
	getName() string
	getAge() int
}

type Player struct {
	name string
	age  int
}

func (player *Player) setName(name string) {}

func (player *Player) setAge(age int) {}

func (player *Player) getName() string {
	return player.name
}

func (player *Player) getAge() int {
	return player.age
}

func getPlayer() (IPlayer, error) {
	return &Player{}, nil
}

func interfacesStruct() {
	var player Player = Player{}
	player.setName("vignesh")
	player.setAge(20)
}
