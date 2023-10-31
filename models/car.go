package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"math/rand"
	"time"
)

type Car struct {
	id          int
	parkingTime time.Duration
	image       *canvas.Image
}

func NewCar(id int) *Car {
	image := canvas.NewImageFromURI(storage.NewFileURI("./assets/car.png"))
	return &Car{
		id:          id,
		parkingTime: time.Duration(rand.Intn(7)+1) * time.Second,
		image:       image,
	}
}

func (c *Car) Enter(p *Parking, carsContainer *fyne.Container) {
	p.GetEntrance().Lock()
	defer p.GetEntrance().Unlock()

	p.GetSpaces() <- c.GetId()
	fmt.Printf("Vehicle %d entered the parking. Spaces left: %d\n", c.GetId(), len(p.GetSpaces()))
	c.image.Move(fyne.NewPos(700, float32(50+(len(p.GetSpaces())*35))))
	carsContainer.Refresh()
}

func (c *Car) Leave(p *Parking, carsContainer *fyne.Container) {
	<-p.GetSpaces()
	fmt.Printf("Vehicle %d left the parking. Spaces left: %d\n", c.GetId(), len(p.GetSpaces()))
	c.image.Move(fyne.NewPos(50, 210))
	carsContainer.Refresh()
}

func (c *Car) Park(p *Parking, carsContainer *fyne.Container) {
	c.image.Resize(fyne.NewSize(50, 30))
	c.image.Move(fyne.NewPos(50, 310))

	carsContainer.Add(c.image)
	carsContainer.Refresh()

	time.Sleep(time.Second)

	c.Enter(p, carsContainer)

	time.Sleep(c.parkingTime)
	c.Leave(p, carsContainer)
}

func (c *Car) GetId() int {
	return c.id
}

func (c *Car) GetCarImage() *canvas.Image {
	return c.image
}
