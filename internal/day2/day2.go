package day2

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func safetyAnalysis(arr []string) (string, error) {
	safe := "safe"
	for i := 1; i < len(arr)-1; i++ {
		left, err := strconv.Atoi(arr[i-1])
		if err != nil {
			return "", nil
		}
		curr, err := strconv.Atoi(arr[i])
		if err != nil {
			return "", nil
		}
		right, err := strconv.Atoi(arr[i+1])
		if err != nil {
			return "", nil
		}

		if (left > curr) && (curr > right) {
			if (left-curr > 3) || (curr-right > 3) {
				safe = "unsafe"
				break
			} else {
				continue
			}
		} else if (left < curr) && (curr < right) {
			if (curr-left > 3) || (right-curr > 3) {
				safe = "unsafe"
				break
			} else {
				continue
			}
		} else {
			safe = "unsafe"
			break
		}
	}
	return safe, nil
}

func ExecutePart1(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	safety := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Fields(text)

		safe, err := safetyAnalysis(split)
		if err != nil {
			return 0, err
		}
		safety = append(safety, safe)
	}

	sum := 0
	for _, elem := range safety {
		if elem == "safe" {
			sum = sum + 1
		}
	}

	return sum, nil
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func ExecutePart2(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	safety := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Fields(text)

		safe, err := safetyAnalysis(split)
		if err != nil {
			return 0, err
		}
		if safe == "unsafe" {
			t_safety := []string{}
			for i := 0; i < len(split); i++ {
				new_copy := make([]string, len(split))
				copy(new_copy, split)
				temp := removeIndex(new_copy, i)
				t_safe, err := safetyAnalysis(temp)
				if err != nil {
					return 0, err
				}
				t_safety = append(t_safety, t_safe)
			}
			if slices.Contains(t_safety, "safe") {
				safe = "safe"
			}
		}
		safety = append(safety, safe)
	}

	sum := 0
	for _, elem := range safety {
		if elem == "safe" {
			sum = sum + 1
		}
	}

	return sum, nil
}
