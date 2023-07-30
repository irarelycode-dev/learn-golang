package main

import (
	"fmt"
)

type ITeams interface {
	makeTeam() ITeam
}

type India struct{}
type Australia struct{}
type England struct{}

func (i *India) makeTeam() ITeam {
	return &IndiaTeam{}
}

func (a *Australia) makeTeam() ITeam {
	return &AustraliaTeam{}
}

func (e *England) makeTeam() ITeam {
	return &EnglandTeam{}
}

type ITeam interface {
	setTeamName(name string)
	getTeamName() string
	addPlayer(name string)
	getPlayers() []string
}

type Team struct {
	name    string
	players []string
}
type IndiaTeam struct {
	Team
}
type AustraliaTeam struct {
	Team
}
type EnglandTeam struct {
	Team
}

func (t *Team) setTeamName(name string) {
	t.name = name
}

func (t *Team) getTeamName() string {
	return t.name
}

func (t *Team) addPlayer(name string) {
	t.players = append(t.players, name)
}

func (t *Team) getPlayers() []string {
	return t.players
}

func abstractFactory(name string) (ITeams, error) {
	if name == "INDIA" {
		return &India{}, nil
	}
	if name == "AUSTRALIA" {
		return &Australia{}, nil
	}
	if name == "ENGLAND" {
		return &England{}, nil
	}
	return nil, fmt.Errorf("Wrong team name %s passed", name)
}

func main() {
	teamFactory, _ := abstractFactory("INDIA")
	team := teamFactory.makeTeam()
	team.setTeamName("Indian team")
	fmt.Print("team name:", team.getTeamName())
}
