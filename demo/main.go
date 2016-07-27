package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/seiflotfy/go-ema"
)

func main() {
	sum := 0.0
	values := make([]float64, 100, 100)

	rand.Seed(time.Now().UnixNano())

	ema1, err := ema.NewExpMovingAverage(10*time.Millisecond, time.Millisecond, 0.0)
	if err != nil {
		panic(err)
	}

	avg := 0.0
	for i := range values {
		time.Sleep(time.Millisecond)
		num := float64(rand.Int() % 100)

		avg = ema1.Add(num)
		if i >= len(values)-10 {
			sum += num
			fmt.Println(">>", i, num, sum, ">>>>>>>>", avg)
		}
	}

	fmt.Println("===========================")
	fmt.Println(avg, "~=", sum/10)
}
