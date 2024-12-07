package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mowshon/iterium"
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

func evaluateMath(row string) (int, error) {
	first := strings.Split(row, ": ")
	target, err := strconv.Atoi(first[0])
	if err != nil {
		return 0, err
	}

	second := strings.Split(first[1], " ")
	vals := []int{}
	for _, str := range second {
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
		vals = append(vals, num)
	}

	opts := []string{"*", "+"}
	combos := iterium.Product(opts, len(second)-1)
	for try := range combos.Chan() {
		total := 0
		for i, x := range try {
			if i == 0 {
				if x == "*" {
					total = vals[0] * vals[i+1]
				} else {
					total = vals[0] + vals[i+1]
				}
			} else {
				if x == "*" {
					total = total * vals[i+1]
				} else {
					total = total + vals[i+1]
				}
			}
		}
		if total == target {
			return target, nil
		}
	}
	return -1, nil
}

func evaluateMathConcat(row string) (int, error) {
	first := strings.Split(row, ": ")
	target, err := strconv.Atoi(first[0])
	if err != nil {
		return 0, err
	}

	second := strings.Split(first[1], " ")
	vals := []int{}
	for _, str := range second {
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
		vals = append(vals, num)
	}

	opts := []string{"*", "+", "||"}
	combos := iterium.Product(opts, len(second)-1)
	for try := range combos.Chan() {
		total := 0
		for i, x := range try {
			if i == 0 {
				if x == "*" {
					total = vals[0] * vals[i+1]
				} else if x == "+" {
					total = vals[0] + vals[i+1]
				} else {
					newNum, err := strconv.Atoi(strconv.Itoa(vals[0]) + strconv.Itoa(vals[i+1]))
					if err != nil {
						return 0, err
					}
					total = newNum
				}
			} else {
				if x == "*" {
					total = total * vals[i+1]
				} else if x == "+" {
					total = total + vals[i+1]
				} else {
					newNum, err := strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(vals[i+1]))
					if err != nil {
						return 0, err
					}
					total = newNum
				}
			}
		}
		if total == target {
			return total, nil
		}
	}
	return -1, nil
}

func ExecutePart1(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		target, err := evaluateMath(text)
		if err != nil {
			return 0, err
		} else if target != -1 {
			total = total + target
		}
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
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		target, err := evaluateMathConcat(text)
		if err != nil {
			return 0, err
		} else if target != -1 {
			total = total + target
		}
	}

	return total, nil
}
