package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

func SecondHandPoint(t time.Time) Point {
	return angletoPoint(SecondsInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / minutesInClock) +
		(math.Pi / (minutesInHalfClock / (float64(t.Minute()))))
}

func minuteHandPoint(t time.Time) Point {
	return angletoPoint(minutesInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) +
		(math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}

func hourHandPoint(t time.Time) Point {
	return angletoPoint(hoursInRadians(t))
}

func angletoPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
