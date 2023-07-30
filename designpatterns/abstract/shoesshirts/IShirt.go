package main

type IShirt interface {
	setLogo(logo string)
	getLogo() string
	setSize(size int)
	getSize() int
}

type Shirt struct {
	logo string
	size int
}

type AdidasShirt struct {
	Shirt
}

type NikeShirt struct {
	Shirt
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getSize() int {
	return s.size
}
