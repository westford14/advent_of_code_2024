package day8

import "testing"

func TestDay8Part1(t *testing.T) {
	got, err := Execute("data/input.txt", 1)
	if err != nil {
		t.Error(err)
	}
	wanted := 14

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}

func TestDay8Part2(t *testing.T) {
	got, err := Execute("data/input.txt", 2)
	if err != nil {
		t.Error(err)
	}
	wanted := 34

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}
