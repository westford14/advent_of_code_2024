package day9

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type File struct {
	ID   int
	Size int
}

func Execute(path string, part int) (int, error) {
	switch part {
	case 1:
		return ExecutePart(path, true)
	case 2:
		return ExecutePart(path, false)
	default:
		return 0, fmt.Errorf("did not understand %d", part)
	}
}

func run(fs []File) (checksum int) {
	for file := len(fs) - 1; file >= 0; file-- {
		for free := 0; free < file; free++ {
			if fs[file].ID != -1 && fs[free].ID == -1 && fs[free].Size >= fs[file].Size {
				fs = slices.Insert(fs, free, fs[file])
				fs[file+1].ID = -1
				fs[free+1].Size = fs[free+1].Size - fs[file+1].Size
			}
		}
	}
	i := 0
	for _, f := range fs {
		for range f.Size {
			if f.ID != -1 {
				checksum += i * f.ID
			}
			i++
		}
	}
	return checksum
}

func ExecutePart(path string, first bool) (int, error) {
	input, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	diskmap := strings.TrimSpace(string(input)) + "0"
	fs1, fs2 := []File{}, []File{}
	for id := 0; id*2 < len(diskmap); id++ {
		size, free := int(diskmap[id*2]-'0'), int(diskmap[id*2+1]-'0')
		fs1 = append(fs1, slices.Repeat([]File{{id, 1}}, size)...)
		fs1 = append(fs1, slices.Repeat([]File{{-1, 1}}, free)...)
		fs2 = append(fs2, File{id, size}, File{-1, free})
	}
	if first {
		return run(fs1), nil
	} else {
		return run(fs2), nil
	}
}
