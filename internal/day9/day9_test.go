package day9

import "testing"

func TestDay9Part1(t *testing.T) {
	got, err := Execute("data/input.txt", 1)
	if err != nil {
		t.Error(err)
	}
	wanted := 1928

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}

func TestDay9Part2(t *testing.T) {
	got, err := Execute("data/input.txt", 2)
	if err != nil {
		t.Error(err)
	}
	wanted := 2858

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}
