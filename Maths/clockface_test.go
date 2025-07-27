package clockface_test

import (
	clockface "learn-golang/Maths"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	expected := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
