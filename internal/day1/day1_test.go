package day1

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	got, err := ExecutePart1(filepath.Join(dir, "data/input.txt"))
	if err != nil {
		t.Error(err.Error())
	}
	wanted := 11
	if got != wanted {
		t.Errorf("%d got but wanted %d", got, wanted)
	}
}

func TestDay1Part2(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	got, err := ExecutePart2(filepath.Join(dir, "data/input.txt"))
	if err != nil {
		t.Error(err.Error())
	}
	wanted := 31
	if got != wanted {
		t.Errorf("%d got but wanted %d", got, wanted)
	}
}
