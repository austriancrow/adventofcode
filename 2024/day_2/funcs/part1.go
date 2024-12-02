package day2

import (
	"math"
	"strconv"
	"strings"
	"time"
	"utilities"
)

func Parse_answer_one(file_content []byte) float64 {
	defer utilities.TimeTrack(time.Now(), "Day 2 Part 1")
	lines := strings.Split(string(file_content), "\n")
	count := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		split_line := strings.Split(line, " ")
		previous_number := -1
		increasing := false
		safe := false
		for index, value := range split_line {
			current_number, _ := strconv.Atoi(value)
			if previous_number == -1 {
				previous_number = current_number
				continue
			} else {
				if index == 1 {
					if current_number > previous_number {
						increasing = true
					} else if current_number < previous_number {
						increasing = false
					} else {
						safe = false
						break
					}
				} else {
					if increasing {
						if current_number > previous_number {
							safe = true
						} else {
							safe = false
							break
						}
					} else {
						if current_number < previous_number {
							safe = true
						} else {
							safe = false
							break
						}
					}
				}
				distance := math.Abs(float64(current_number - previous_number))
				if distance >= 1 && distance <= 3 {
					safe = true
				} else {
					safe = false
					break
				}
				previous_number = current_number
			}
		}
		if safe {
			count++
		}
	}

	return float64(count)
}
