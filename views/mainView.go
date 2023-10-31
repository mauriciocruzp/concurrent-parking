package views

import (
	"concurrent-parking/models"
	"concurrent-parking/scenes"
	"fmt"
	"sync"
	"time"
)

type MainView struct{}

func NewMainView() *MainView {
	return &MainView{}
}

func (v *MainView) Show() {

}

func (v *MainView) Run() {
	p := scenes.NewParking(make(chan int, 20), &sync.Mutex{})
	poissonDist := models.NewPoissonDist()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			v := models.NewCar(id)
			v.Park(p)
		}(i)
		var randPoissonNumber = poissonDist.Generate(float64(3))
		fmt.Println(randPoissonNumber)
		time.Sleep(time.Second * time.Duration(randPoissonNumber))
	}

	wg.Wait()
}
