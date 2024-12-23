package day7

import "testing"

func TestDay7Part1(t *testing.T) {
	got, err := Execute("data/input.txt", 1)
	if err != nil {
		t.Error(err)
	}
	wanted := 3749

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}

func TestDay7Part2(t *testing.T) {
	got, err := Execute("data/input.txt", 2)
	if err != nil {
		t.Error(err)
	}
	wanted := 11387

	if got != wanted {
		t.Errorf("got %d, but wanted %d", got, wanted)
	}
}
