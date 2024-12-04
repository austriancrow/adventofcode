package day4

import (
	"time"
	"utilities"
)

func search_deeper2(current_field *Field, val0 string, val1 string) bool {
	if current_field.neighbors[1] == nil ||
		current_field.neighbors[3] == nil ||
		current_field.neighbors[5] == nil ||
		current_field.neighbors[7] == nil {
		return false
	}
	value_ne := current_field.neighbors[1].value
	value_se := current_field.neighbors[3].value
	value_sw := current_field.neighbors[5].value
	value_nw := current_field.neighbors[7].value
	if ((value_ne == val0 && value_sw == val1) ||
		(value_ne == val1 && value_sw == val0)) &&
		((value_se == val0 && value_nw == val1) ||
			(value_se == val1 && value_nw == val0)) {
		return true
	}

	return false
}

func Parse_answer_two(file_content []byte) float64 {
	defer utilities.TimeTrack(time.Now(), "Day 4 Part 2")
	// lines := strings.Split(string(file_content), "\n")
	total := 0
	fieldMap := NewMap(file_content)
	for _, row := range fieldMap.fields {
		for _, field := range row {
			if field.value == "A" {
				answer := search_deeper2(field, "M", "S")
				if answer {
					total++
				}
			}
		}
	}

	return float64(total)
}
