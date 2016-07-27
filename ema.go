package ema

import (
	"math"
	"time"
)

// ExpMovingAverage ...
type ExpMovingAverage struct {
	avg                float64
	ttl                time.Duration
	tick               time.Duration
	lastTimestamp      time.Time
	lastFlushTimestamp time.Time
	buffer             float64
}

// NewExpMovingAverage ...
func NewExpMovingAverage(ttl, tick time.Duration, initValue float64) (*ExpMovingAverage, error) {
	return &ExpMovingAverage{
		avg:                initValue,
		ttl:                ttl,
		tick:               tick,
		lastTimestamp:      time.Now(),
		lastFlushTimestamp: time.Now(),
	}, nil
}

func (ema *ExpMovingAverage) add(value float64, delta int64) float64 {
	ed := float64(delta) / float64(ema.ttl.Nanoseconds())
	alpha := 1.0 - math.Exp(-ed)
	return (1.0-alpha)*ema.avg + alpha*value
}

// Add ...
func (ema *ExpMovingAverage) Add(value float64) float64 {
	delta := time.Since(ema.lastTimestamp)
	ema.lastTimestamp = time.Now()
	ema.avg = ema.add(value, delta.Nanoseconds())
	return ema.avg
}
