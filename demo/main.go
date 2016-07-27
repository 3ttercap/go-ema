package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/seiflotfy/go-ema"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	sum := 0.0
	duration := 1000 * time.Millisecond
	windowSize := 10.0

	avg := 0.0
	// keep track of last 10ms
	ema1, err := ema.NewExpMovingAverage(time.Duration(windowSize)*time.Millisecond, 0.0)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	// run for almost 15 ms

	for time.Since(now) < duration {
		time.Sleep(time.Millisecond)
		num := float64(rand.Int() % 100)
		avg = ema1.Add(num)
		if duration-time.Since(now) < time.Duration(windowSize)*time.Millisecond {
			sum += num
		}
	}

	fmt.Println(avg, "~=", sum/windowSize)

	//////

	fmt.Println("============")
	em, err := ema.NewExpMovingAverage(10*time.Millisecond, 0.0)

	for i := 1.0; i <= 100; i++ {
		time.Sleep(1 * time.Millisecond)
		avg = em.Add(i)
	}

	fmt.Println(em.Get())
}
