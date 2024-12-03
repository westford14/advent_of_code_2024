package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	mul_match, err := regexp.Compile("mul\\(\\d+,\\d+\\)")
	if err != nil {
		return 0, err
	}
	int_match, err := regexp.Compile("\\d+")
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		matched := mul_match.FindAllStringSubmatch(text, -1)

		for _, elem := range matched {
			ints := int_match.FindAllStringSubmatch(elem[0], -1)
			int_1, err := strconv.Atoi(ints[0][0])
			if err != nil {
				return 0, err
			}
			int_2, err := strconv.Atoi(ints[1][0])
			if err != nil {
				return 0, err
			}
			total = total + (int_1 * int_2)
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

	mul_match, err := regexp.Compile("mul\\(\\d+,\\d+\\)")
	if err != nil {
		return 0, err
	}
	int_match, err := regexp.Compile("\\d+")
	if err != nil {
		return 0, err
	}
	op_match, err := regexp.Compile("don't\\(\\)|do\\(\\)")
	if err != nil {
		return 0, err
	}

	full_string := ""
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		full_string = full_string + text
	}

	copied_text := strings.Clone(full_string)
	do := true
	for {
		op_loc := op_match.FindStringIndex(copied_text)
		mul_loc := mul_match.FindStringIndex(copied_text)
		if mul_loc == nil {
			break
		}
		if op_loc == nil {
			if do {
				substring := copied_text[mul_loc[0]:mul_loc[1]]
				matched := mul_match.FindAllStringSubmatch(substring, -1)

				for _, elem := range matched {
					ints := int_match.FindAllStringSubmatch(elem[0], -1)
					int_1, err := strconv.Atoi(ints[0][0])
					if err != nil {
						return 0, err
					}
					int_2, err := strconv.Atoi(ints[1][0])
					if err != nil {
						return 0, err
					}
					total = total + (int_1 * int_2)
				}
				copied_text = copied_text[mul_loc[1]:]
			} else {
				copied_text = copied_text[mul_loc[1]:]
			}
			continue
		}
		if op_loc[0] < mul_loc[0] {
			if copied_text[op_loc[0]:op_loc[1]] == "don't()" {
				do = false
				copied_text = copied_text[op_loc[1]:]
			} else if copied_text[op_loc[0]:op_loc[1]] == "do()" {
				do = true
				copied_text = copied_text[op_loc[1]:]
			}
		} else {
			if do {
				substring := copied_text[mul_loc[0]:mul_loc[1]]
				matched := mul_match.FindAllStringSubmatch(substring, -1)

				for _, elem := range matched {
					ints := int_match.FindAllStringSubmatch(elem[0], -1)
					int_1, err := strconv.Atoi(ints[0][0])
					if err != nil {
						return 0, err
					}
					int_2, err := strconv.Atoi(ints[1][0])
					if err != nil {
						return 0, err
					}
					total = total + (int_1 * int_2)
				}
				copied_text = copied_text[mul_loc[1]:]
			} else {
				copied_text = copied_text[mul_loc[1]:]
			}
		}
	}
	return total, nil
}
