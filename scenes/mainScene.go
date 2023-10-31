package scenes

import (
	"concurrent-parking/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"sync"
	"time"
)

type MainScene struct {
	window fyne.Window
}

func NewMainScene(window fyne.Window) *MainScene {
	return &MainScene{
		window: window,
	}
}

var carsContainer = container.NewWithoutLayout()

func (s *MainScene) Show() {
	s.window.SetContent(carsContainer)
}

func (s *MainScene) Run() {
	s.Show()
	p := models.NewParking(make(chan int, 20), &sync.Mutex{})
	poissonDist := models.NewPoissonDist()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			car := models.NewCar(id)

			car.Park(p, carsContainer)
		}(i)
		var randPoissonNumber = poissonDist.Generate(float64(2))
		fmt.Println(randPoissonNumber)
		time.Sleep(time.Second * time.Duration(randPoissonNumber))
	}

	wg.Wait()
}
