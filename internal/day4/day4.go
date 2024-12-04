package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var directions = map[string][]int{
	"UP-LEFT":    {-1, -1},
	"UP-MID":     {-1, 0},
	"UP-RIGHT":   {-1, 1},
	"LEFT-MID":   {0, -1},
	"MID-RIGHT":  {0, 1},
	"DOWN-LEFT":  {1, -1},
	"DOWN-MID":   {1, 0},
	"DOWN-RIGHT": {1, 1},
}
var crossDirections = map[string][]int{
	"UP-LEFT":    {-1, -1},
	"UP-RIGHT":   {-1, 1},
	"DOWN-LEFT":  {1, -1},
	"DOWN-RIGHT": {1, 1},
}
var reverseDirections = map[string][]int{
	"UP-LEFT":    {1, 1},
	"UP-RIGHT":   {1, -1},
	"DOWN-LEFT":  {-1, 1},
	"DOWN-RIGHT": {-1, -1},
}

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

func searchGrid(matrix [][]string, i, j, i_lim, j_lim int, target, direction string) (bool, []int, string) {
	x := i + directions[direction][0]
	y := j + directions[direction][1]
	if x < 0 || y < 0 {
		return false, nil, ""
	} else if x >= i_lim || y >= j_lim {
		return false, nil, ""
	} else {
		if matrix[x][y] == target {
			return true, []int{x, y}, direction
		} else {
			return false, nil, ""
		}
	}
}

func traverseGrid(matrix [][]string) int {
	i_lim := len(matrix)
	j_lim := len(matrix[0])
	total := 0
	for i, row := range matrix {
		for j, letter := range row {
			if letter == "X" {
				for dir, _ := range directions {
					found, index, direction := searchGrid(matrix, i, j, i_lim, j_lim, "M", dir)
					if !found {
						continue
					} else {
						a_found, a_index, a_direction := searchGrid(matrix, index[0], index[1], i_lim, j_lim, "A", direction)
						if !a_found {
							continue
						} else {
							s_found, _, _ := searchGrid(matrix, a_index[0], a_index[1], i_lim, j_lim, "S", a_direction)
							if !s_found {
								continue
							} else {
								total = total + 1
							}
						}
					}
				}
			}
		}
	}
	return total
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

	return traverseGrid(matrix), nil
}

func crossGrid(matrix [][]string) int {
	i_lim := len(matrix)
	j_lim := len(matrix[0])
	total := 0
	for i, row := range matrix {
		for j, letter := range row {
			if letter == "A" {
				combos := [][]string{
					{"UP-LEFT", "UP-RIGHT"},
					{"UP-RIGHT", "DOWN-RIGHT"},
					{"DOWN-RIGHT", "DOWN-LEFT"},
					{"DOWN-LEFT", "UP-LEFT"},
				}
				for _, dir := range combos {
					first := crossDirections[dir[0]]
					second := crossDirections[dir[1]]

					x := i + first[0]
					y := j + first[1]
					x_2 := i + second[0]
					y_2 := j + second[1]
					if x < 0 || y < 0 {
						continue
					} else if x_2 < 0 || y_2 < 0 {
						continue
					} else if x_2 >= i_lim || y_2 >= j_lim {
						continue
					} else if x >= i_lim || y >= j_lim {
						continue
					} else {
						if matrix[x][y] == "M" && matrix[x_2][y_2] == "M" {
							first_s := reverseDirections[dir[0]]
							second_s := reverseDirections[dir[1]]

							x_s := i + first_s[0]
							y_s := j + first_s[1]
							x_s_2 := i + second_s[0]
							y_s_2 := j + second_s[1]

							if x_s < 0 || y_s < 0 {
								continue
							} else if x_s_2 < 0 || y_s_2 < 0 {
								continue
							} else if x_s_2 >= i_lim || y_s_2 >= j_lim {
								continue
							} else if x_s >= i_lim || y_s >= j_lim {
								continue
							} else {
								if matrix[x_s][y_s] == "S" && matrix[x_s_2][y_s_2] == "S" {
									total = total + 1
								}
							}
						} else {
							continue
						}
					}
				}
			}
		}
	}
	return total
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
	return crossGrid(matrix), nil
}
