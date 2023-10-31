package models

import (
	"sync"
)

type Parking struct {
	spaces      chan int
	entrance    *sync.Mutex
	spacesArray [20]bool
}

func NewParking(spaces chan int, entrance *sync.Mutex) *Parking {
	return &Parking{
		spaces:      spaces,
		entrance:    entrance,
		spacesArray: [20]bool{},
	}
}

func (p *Parking) GetSpaces() chan int {
	return p.spaces
}

func (p *Parking) GetEntrance() *sync.Mutex {
	return p.entrance
}
