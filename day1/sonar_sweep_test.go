package day1

import "testing"

func TestCountDepthIncreased(t *testing.T) {
	got := CountDepthIncreased([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	want := 7

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
