package day2

import "testing"

func TestDay2Part1(t *testing.T) {
	got, err := Execute("data/input.txt", "part1")
	if err != nil {
		t.Error(err)
	}
	wanted := 2

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}

func TestDay2Part2(t *testing.T) {
	got, err := Execute("data/input.txt", "part2")
	if err != nil {
		t.Error(err)
	}
	wanted := 4

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}
