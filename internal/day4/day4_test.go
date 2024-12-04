package day4

import "testing"

func TestDay4Part1(t *testing.T) {
	got, err := Execute("data/input.txt", 1)
	if err != nil {
		t.Error(err)
	}
	wanted := 18

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}

func TestDay4Part2(t *testing.T) {
	got, err := Execute("data/input.txt", 2)
	if err != nil {
		t.Error(err)
	}
	wanted := 9

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}
