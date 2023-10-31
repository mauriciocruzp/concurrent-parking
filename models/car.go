package models

import (
	"concurrent-parking/scenes"
	"fmt"
	"math/rand"
	"time"
)

type Car struct {
	id          int
	parkingTime time.Duration
}

func NewCar(id int) *Car {
	return &Car{
		id:          id,
		parkingTime: time.Duration(rand.Intn(5)+1) * time.Second,
	}
}

func (c *Car) Enter(p *scenes.Parking) {
	p.GetEntrance().Lock()
	defer p.GetEntrance().Unlock()

	p.Spaces <- c.GetId()
	fmt.Printf("Vehicle %d entered the parking. Spaces left: %d\n", c.GetId(), len(p.Spaces))
}

func (c *Car) Leave(p *scenes.Parking) {
	<-p.GetSpaces()
	fmt.Printf("Vehicle %d left the parking. Spaces left: %d\n", c.GetId(), len(p.Spaces))
}

func (c *Car) Park(p *scenes.Parking) {
	c.Enter(p)
	time.Sleep(c.parkingTime)
	c.Leave(p)
}

func (c *Car) GetId() int {
	return c.id
}
