package day3

import (
	"regexp"
	"strconv"
	"strings"
	"time"
	"utilities"
)

func Parse_answer_one(file_content []byte) float64 {
	defer utilities.TimeTrack(time.Now(), "Day 3 Part 1")
	lines := strings.Split(string(file_content), "\n")
	total := 0
	regex_mul_string, _ := regexp.Compile(`mul\((?P<num0>\d+),(?P<num1>\d+)\)`)

	for _, line := range lines {
		if line == "" {
			continue
		}
		matches := regex_mul_string.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			num0, _ := strconv.Atoi(match[1])
			num1, _ := strconv.Atoi(match[2])
			total += num0 * num1

		}
	}

	return float64(total)
}
