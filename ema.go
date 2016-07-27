package ema

import (
	"math"
	"time"
)

// ExpMovingAverage ...
type ExpMovingAverage struct {
	avg           float64
	ttl           time.Duration
	tick          time.Duration
	lastTimestamp time.Time
	buffer        float64
	bufferSize    float64
}

// NewExpMovingAverage ...
func NewExpMovingAverage(ttl time.Duration, initValue float64) (*ExpMovingAverage, error) {
	return &ExpMovingAverage{
		avg:           initValue,
		ttl:           ttl,
		tick:          ttl / 1000,
		lastTimestamp: time.Now(),
	}, nil
}

func (ema *ExpMovingAverage) add(value, avg float64, delta int64) float64 {
	offset := 1.0
	ed := float64(delta) / float64(ema.ttl.Nanoseconds())
	alpha := offset - math.Exp(-ed)
	return (offset-alpha)*avg + alpha*value
}

// Add adds a value and returns the new average
func (ema *ExpMovingAverage) Add(value float64) float64 {
	now := time.Now()
	delta := time.Since(ema.lastTimestamp)
	// if delta is greater than tick then flush the buffer
	if delta > ema.tick {
		// flush the buffer
		bValue := 0.0
		if ema.buffer > 0 && ema.bufferSize > 0 {
			bValue = ema.buffer / ema.bufferSize
		}
		ema.avg = ema.add(bValue, ema.avg, delta.Nanoseconds())
		ema.buffer = 0
		ema.bufferSize = 0
	}
	ema.lastTimestamp = now
	if value > 0 {
		ema.buffer += value
		ema.bufferSize++
	}

	bValue := 0.0
	if ema.buffer > 0 && ema.bufferSize > 0 {
		bValue = ema.buffer / ema.bufferSize
	}

	return ema.add(bValue, ema.avg, delta.Nanoseconds())
}

// Get returns the current average
func (ema *ExpMovingAverage) Get() float64 {
	// Add 0.0 to make sure we decay when requested
	return ema.Add(0.0)
}
