package views

import (
	"concurrent-parking/models"
	"concurrent-parking/scenes"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MainView struct{}

func NewMainView() *MainView {
	return &MainView{}
}

func (v *MainView) Run() {
	rand.Seed(time.Now().UnixNano())

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
		var randPoissonNumber = poissonDist.Generate(float64(143))
		fmt.Println(randPoissonNumber)
		time.Sleep(time.Second * time.Duration(randPoissonNumber))
	}

	wg.Wait()
}
