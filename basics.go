package main

import (
	"fmt"
	"time"
)

func fun(value string) {
	for i := 0; i < 3; i++ {
		fmt.Println(value)
		time.Sleep(1 * time.Millisecond)
	}
}

func mainFun() {
	// Direct call
	fun("Direct call")

	// write goroutines with differents variants for function call
	fgx := fun
	go fgx("goroutine - 2")
	// goroutine function call
	go fun("Goroutine - 1") // não chega a printar pois não espera, é tratado com channel ou com sleep

	// goroutines with anonymous value call
	go func() {
		fun("goroutine - 3")
	}()
	// wait for goroutine to end
	time.Sleep(5 * time.Millisecond)
	fmt.Println("Done")
}
