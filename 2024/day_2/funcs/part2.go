package day2

import (
	"math"
	"strconv"
	"strings"
	"time"
	"utilities"
)

func check_increasing(num0 int, num1 int) bool {
	if num0 < num1 {
		return true
	}
	return false
}

func check_decreasing(num0 int, num1 int) bool {
	if num0 > num1 {
		return true
	}
	return false
}

func check_distance(num0 int, num1 int) bool {
	distance := math.Abs(float64(num0 - num1))
	if distance >= 1 && distance <= 3 {
		return true
	}
	return false
}

func check_safe(num0 int, num1 int, increasing bool) bool {
	safe := false
	if increasing {
		safe = check_increasing(num0, num1)
	} else {
		safe = check_decreasing(num0, num1)
	}
	if !safe {
		return safe
	}
	safe = check_distance(num0, num1)
	return safe
}

type pair struct {
	num0 int
	num1 int
}

func Parse_answer_two(file_content []byte) float64 {
	defer utilities.TimeTrack(time.Now(), "Day 2 Part 2")
	lines := strings.Split(string(file_content), "\n")
	count := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		split_line := strings.Split(line, " ")
		var number_array []int
		var pair_array []pair
		for index, value := range split_line {
			current_number, _ := strconv.Atoi(value)
			number_array = append(number_array, current_number)
			if index > 0 {
				pair_array = append(pair_array, pair{number_array[index-1], current_number})
			}
		}
		permutations := [][]pair{pair_array}
		for index, _ := range number_array {
			var pair_array []pair
			previous_number := -1
			for iindex, value := range number_array {
				if iindex == index {
					continue
				}
				if previous_number == -1 {
					previous_number = value
					continue
				}
				pair_array = append(pair_array, pair{previous_number, value})
				previous_number = value
			}
			permutations = append(permutations, pair_array)
		}
		safe := false
		for _, pair_array := range permutations {
			increasing := pair_array[0].num0 < pair_array[0].num1
			for _, pair := range pair_array {
				safe = check_safe(pair.num0, pair.num1, increasing)
				if !safe {
					break
				}
			}
			if safe {
				break
			}
		}

		if safe {
			count++
		}

	}

	return float64(count)
}
