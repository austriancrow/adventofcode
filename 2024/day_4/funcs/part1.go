package day4

import (
	"time"
	"utilities"
)

func search_deeper(current_field *Field, search_array []string, search_index int, search_direction int) bool {
	if current_field.value == search_array[search_index] {
		if search_index == len(search_array)-1 {
			return true
		}
		if current_field.neighbors[search_direction] != nil {
			return search_deeper(current_field.neighbors[search_direction], search_array, search_index+1, search_direction)
		}
	}
	return false
}

func Parse_answer_one(file_content []byte) float64 {
	defer utilities.TimeTrack(time.Now(), "Day 3 Part 1")
	// lines := strings.Split(string(file_content), "\n")
	total := 0
	fieldMap := NewMap(file_content)
	search_array := []string{"X", "M", "A", "S"}
	for _, row := range fieldMap.fields {
		for _, field := range row {
			if field.value == search_array[0] {
				for i := 0; i < 8; i++ {
					answer := search_deeper(&field, search_array, 0, i)
					if answer {
						total++
					}
				}
			}
		}
	}

	return float64(total)
}
