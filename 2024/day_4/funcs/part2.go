package day4

import (
	"regexp"
	"strconv"
	"strings"
	"time"
	"utilities"
)

func Parse_answer_two(file_content []byte) float64 {
	defer utilities.TimeTrack(time.Now(), "Day 3 Part 2")
	lines := strings.Split(string(file_content), "\n")
	total := 0
	regex_all_commands, _ := regexp.Compile(`mul\((?P<num0>\d+),(?P<num1>\d+)\)|don't\(\)|do\(\)`)
	dont := false
	for _, line := range lines {
		if line == "" {
			continue
		}
		matches := regex_all_commands.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			command := match[0]
			if command == "don't()" {
				dont = true
				continue
			}
			if command == "do()" {
				dont = false
				continue
			}
			if !dont {
				num0, _ := strconv.Atoi(match[1])
				num1, _ := strconv.Atoi(match[2])
				total += num0 * num1
			}

		}
	}
	return float64(total)
}
