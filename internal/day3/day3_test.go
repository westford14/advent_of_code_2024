package day3

import "testing"

func TestDay3Part1(t *testing.T) {
	got, err := Execute("data/input.txt", "part1")
	if err != nil {
		t.Error(err)
	}
	wanted := 161

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}

func TestDay3Part2(t *testing.T) {
	got, err := Execute("data/input_2.txt", "part2")
	if err != nil {
		t.Error(err)
	}
	wanted := 48

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}
