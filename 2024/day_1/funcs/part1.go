package day1

import (
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
	"utilities"
)

func Parse_answer_one(file_content []byte) float64 {
	defer utilities.TimeTrack(time.Now(), "Day 1 Part 1")
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
	total_distance := 0.0
	for index, value_left := range left {
		value_right := right[index]
		total_distance += math.Abs(value_right - value_left)
	}

	return total_distance
}
