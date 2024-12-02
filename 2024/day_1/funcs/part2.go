package day1

import (
	"slices"
	"strconv"
	"strings"
	"time"
	"utilities"
)

func Parse_answer_two(file_content []byte) float64 {
	defer utilities.TimeTrack(time.Now(), "Day 1 Part 2")
	lines := strings.Split(string(file_content), "\n")
	var left []float64
	var right []float64
	for _, line := range lines {
		split_line := strings.Split(line, "   ")
		value0, _ := strconv.Atoi(split_line[0])
		value1, _ := strconv.Atoi(split_line[1])
		left = append(left, float64(value0))
		right = append(right, float64(value1))
	}
	slices.Sort(left)
	slices.Sort(right)
	right_map := make(map[float64]float64)
	for _, value := range right {
		if _, ok := right_map[value]; !ok {
			right_map[value] = 0
		}
		right_map[value] += 1.0
	}
	total_distance := 0.0
	for _, value_left := range left {
		value_right := right_map[value_left]
		total_distance += value_right * value_left
	}

	return total_distance
}
