package scenes

import (
	"sync"
)

type Parking struct {
	Spaces   chan int
	Entrance *sync.Mutex
}

func NewParking(spaces chan int, entrance *sync.Mutex) *Parking {
	return &Parking{
		Spaces:   spaces,
		Entrance: entrance,
	}
}

func (p *Parking) GetSpaces() chan int {
	return p.Spaces
}

func (p *Parking) GetEntrance() *sync.Mutex {
	return p.Entrance
}
