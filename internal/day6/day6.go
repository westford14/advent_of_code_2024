package day6

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func Execute(path string, part int) (int, error) {
	switch part {
	case 1:
		return ExecutePart1(path)
	case 2:
		return ExecutePart2(path)
	default:
		return 0, fmt.Errorf("did not understand %d", part)
	}
}

func nextPosition(matrix [][]string, currPosition []int, direction string) ([]int, error) {
	x := currPosition[0]
	y := currPosition[1]
	xLim := len(matrix[0])
	yLim := len(matrix)
	switch direction {
	case "left":
		y = y - 1
		if y < 0 {
			return []int{-1, -1}, nil
		} else {
			return []int{x, y}, nil
		}
	case "right":
		y = y + 1
		if y >= yLim {
			return []int{-1, -1}, nil
		} else {
			return []int{x, y}, nil
		}
	case "up":
		{
			x = x - 1
			if x < 0 {
				return []int{-1, -1}, nil
			} else {
				return []int{x, y}, nil
			}
		}
	case "down":
		{
			x = x + 1
			if x >= xLim {
				return []int{-1, -1}, nil
			} else {
				return []int{x, y}, nil
			}
		}
	default:
		return nil, fmt.Errorf("direction not understood")
	}
}

func plotPath(matrix [][]string) (int, error) {
	var startingPosition = make([]int, 2)
	for row, rElem := range matrix {
		for col, cElem := range rElem {
			if cElem == "^" {
				startingPosition = []int{row, col}
			}
		}
	}

	var currPosition = make([]int, 2)
	copy(currPosition, startingPosition)
	directions := []string{"up", "right", "down", "left"}
	dirInd := 0
	direction := directions[dirInd]
	for {
		next, err := nextPosition(matrix, currPosition, direction)
		if err != nil {
			return 0, err
		} else if slices.Equal([]int{-1, -1}, next) {
			matrix[currPosition[0]][currPosition[1]] = "|"
			break
		} else {
			if matrix[next[0]][next[1]] == "#" {
				dirInd = dirInd + 1
				if dirInd >= len(directions) {
					dirInd = 0
				}
				direction = directions[dirInd]
				continue
			}
			matrix[currPosition[0]][currPosition[1]] = "|"
			currPosition = next
		}
	}

	count := 0
	for _, rElem := range matrix {
		for _, cElem := range rElem {
			if cElem == "|" {
				count = count + 1
			}
		}
	}
	return count, nil
}

func escapePath(matrix [][]string, r, c int) (bool, error) {
	matrix[r][c] = "#"
	var startingPosition = make([]int, 2)
	for row, rElem := range matrix {
		for col, cElem := range rElem {
			if cElem == "^" {
				startingPosition = []int{row, col}
			}
		}
	}

	var currPosition = make([]int, 2)
	copy(currPosition, startingPosition)
	directions := []string{"up", "right", "down", "left"}
	dirInd := 0
	direction := directions[dirInd]
	total := len(matrix) * len(matrix[0])
	count := 1
	for {
		next, err := nextPosition(matrix, currPosition, direction)
		if err != nil {
			return false, err
		} else if slices.Equal([]int{-1, -1}, next) {
			return false, nil
		} else {
			if matrix[next[0]][next[1]] == "#" {
				dirInd = dirInd + 1
				if dirInd >= len(directions) {
					dirInd = 0
				}
				direction = directions[dirInd]
				continue
			}
			count = count + 1
			currPosition = next
			if count >= total {
				return true, nil
			}
		}
	}
}

func ExecutePart1(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	matrix := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, "")
		matrix = append(matrix, split)
	}

	total, err := plotPath(matrix)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func ExecutePart2(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	matrix := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, "")
		matrix = append(matrix, split)
	}

	count := 0
	for r, rElem := range matrix {
		for c, cElem := range rElem {
			if cElem == "#" || cElem == "^" {
				continue
			} else {
				loop, err := escapePath(matrix, r, c)
				if err != nil {
					return 0, err
				}
				if loop {
					count = count + 1
				}
				matrix[r][c] = "."
				continue
			}
		}
	}

	return count, nil
}
