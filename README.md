# Exponential Moving Average

This is meant for sliding windows where the window size is time based and not bucket based
It is self cleansing, since it keeps a buffer is applied (without storing) between ticks

# Usage

```go
// create exponential moving average with the size of 10ms
em, err := ema.NewExpMovingAverage(10*time.Millisecond, 0.0)

for i := 1.0; i <= 100; i++ {
    time.Sleep(1 * time.Millisecond)
    // add i every 1ms
    avg = em.Add(i)
}

em.Get() // ==> ~92
```
