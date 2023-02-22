package clockface

import (
	"testing"
	"time"
)

func TestSecondHangAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	clockface := Clockface{}

	want := clockface.SetPoint(150, 150-90)
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
