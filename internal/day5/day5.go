package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
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

func parseRules(rules []string) (map[int][]int, error) {
	parsedRules := map[int][]int{}
	for _, rule := range rules {
		tokens := strings.Split(rule, "|")
		key, err := strconv.Atoi(tokens[0])
		if err != nil {
			return nil, err
		}
		value, err := strconv.Atoi(tokens[1])
		if err != nil {
			return nil, err
		}

		if _, ok := parsedRules[key]; ok {
			parsedRules[key] = append(parsedRules[key], value)
		} else {
			parsedRules[key] = []int{value}
		}
	}
	return parsedRules, nil
}

func checkOrder(parsedOrder []int, rules map[int][]int) bool {
	correct := true
	orderLen := len(parsedOrder)
	for i := 0; i < orderLen; i++ {
		pageNum := parsedOrder[i]
		pageOrder, ok := rules[pageNum]
		if !ok {
			continue
		}
		if i+1 > orderLen {
			correct = true
		} else {
			for j := i + 1; j < orderLen; j++ {
				testPage := parsedOrder[j]
				if slices.Contains(pageOrder, testPage) {
					continue
				} else {
					correct = false
					return correct
				}
			}
		}
		leftSlice := parsedOrder[:i]
		if len(leftSlice) == 0 {
			continue
		} else {
			for _, lVal := range leftSlice {
				if slices.Contains(pageOrder, lVal) {
					correct = false
					return correct
				}
			}
		}
	}
	return correct
}

func parseOrders(order string, rules map[int][]int) (int, error) {
	parsedOrder := []int{}
	vals := strings.Split(order, ",")
	for _, val := range vals {
		v, err := strconv.Atoi(val)
		if err != nil {
			return 0, err
		}
		parsedOrder = append(parsedOrder, v)
	}

	correct := checkOrder(parsedOrder, rules)

	if correct {
		return parsedOrder[len(parsedOrder)/2], nil
	}
	return -1, nil
}

func ExecutePart1(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	rules := []string{}
	orders := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "|") {
			rules = append(rules, text)
		} else if strings.Contains(text, ",") {
			orders = append(orders, text)
		} else {
			continue
		}
	}
	parsedRules, err := parseRules(rules)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, order := range orders {
		val, err := parseOrders(order, parsedRules)
		if err != nil {
			return 0, err
		}
		if val > 0 {
			total = total + val
		}
	}

	return total, nil
}

func ExecutePart2(path string) (int, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	split := strings.Split(strings.TrimSpace(string(file)), "\n\n")

	compare := func(left, right string) int {
		for _, s := range strings.Split(split[0], "\n") {
			if s := strings.Split(s, "|"); s[0] == left && s[1] == right {
				return -1
			}
		}
		return 0
	}

	run := func(sorted bool) (r int) {
		for _, s := range strings.Split(split[1], "\n") {
			if s := strings.Split(s, ","); slices.IsSortedFunc(s, compare) == sorted {
				slices.SortFunc(s, compare)
				n, _ := strconv.Atoi(s[len(s)/2])
				r += n
			}
		}
		return r
	}

	return run(false), nil
}
