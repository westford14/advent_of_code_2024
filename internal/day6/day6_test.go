package day6

import "testing"

func TestDay6Part1(t *testing.T) {
	got, err := Execute("data/input.txt", 1)
	if err != nil {
		t.Error(err)
	}
	wanted := 41

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}

func TestDay6Part2(t *testing.T) {
	got, err := Execute("data/input.txt", 2)
	if err != nil {
		t.Error(err)
	}
	wanted := 6

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}
