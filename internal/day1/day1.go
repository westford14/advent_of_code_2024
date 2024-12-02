package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Execute(path, part string) (int, error) {
	switch part {
	case "part1":
		return ExecutePart1(path)
	case "part2":
		return ExecutePart2(path)
	default:
		return 0, fmt.Errorf("did not understand %s", part)
	}
}

func ExecutePart1(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	left := []int{}
	right := []int{}
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Fields(text)
		left_int, err := strconv.Atoi(split[0])
		if err != nil {
			return 0, err
		}
		right_int, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, err
		}
		left = append(left, left_int)
		right = append(right, right_int)
	}

	sort.Ints(left)
	sort.Ints(right)

	diff := []int{}
	for i, elem := range left {
		diff = append(diff, int(math.Abs(float64(elem-right[i]))))
	}

	total := 0
	for _, element := range diff {
		total = total + element
	}
	return total, nil
}

func ExecutePart2(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	left := []int{}
	right := []int{}
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Fields(text)
		left_int, err := strconv.Atoi(split[0])
		if err != nil {
			return 0, err
		}
		right_int, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, err
		}
		left = append(left, left_int)
		right = append(right, right_int)
	}

	sim_score := 0
	for _, elem := range left {
		count := 0
		for _, r_elem := range right {
			if elem == r_elem {
				count = count + 1
			}
		}
		sim_score = sim_score + (elem * count)
	}
	return sim_score, nil
}
