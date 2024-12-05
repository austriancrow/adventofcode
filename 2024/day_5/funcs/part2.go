package day5

import (
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
	"utilities"
)

func Parse_answer_two(file_content []byte) float64 {
	defer utilities.TimeTrack(time.Now(), "Day 5 Part 2")
	sections := strings.Split(string(file_content), "\n\n")
	order_lines := strings.Split(sections[0], "\n")
	update_lines := strings.Split(sections[1], "\n")
	total := 0
	order_dict := map[int][]int{}
	for _, order_line := range order_lines {
		if order_line == "" {
			continue
		}
		order := strings.Split(order_line, "|")
		order_key, _ := strconv.Atoi(order[0])
		order_value, _ := strconv.Atoi(order[1])
		if val, ok := order_dict[order_key]; !ok {
			order_dict[order_key] = []int{order_value}
		} else {
			order_dict[order_key] = append(val, order_value)
		}
	}
	for _, update_line := range update_lines {
		if update_line == "" {
			continue
		}
		update_numbers := strings.Split(update_line, ",")
		middle_index := math.Ceil(float64(len(update_numbers)-1) / 2.0)
		// fmt.Println("Middle index", middle_index)
		middle_value := -1
		wrong := true
		// fmt.Println("--------------------")
		// fmt.Println("Checking", update_line)
		new_update_positions := make([]int, len(update_numbers))

		for index, update_value := range update_numbers {
			current_page, _ := strconv.Atoi(update_value)

			for offset, value := range update_numbers[index+1:] {
				next_page, _ := strconv.Atoi(value)
				back_pages := order_dict[next_page]
				if slices.Contains(back_pages, current_page) {
					// fmt.Println(current_page, "should be after", next_page)
					wrong = false
					new_update_positions[index] += 1
					new_update_positions[index+offset+1] -= 1
				}
				// fmt.Println(current_page, "ambivalent about", next_page)
			}
		}
		if !wrong {
			new_update_values := make([]string, len(update_numbers))
			for index, offset := range new_update_positions {
				new_update_values[index+offset] = update_numbers[index]
			}
			// fmt.Println("--------------------")
			// fmt.Println("Old update values", update_numbers)
			// fmt.Println("New update values", new_update_values)
			middle_value, _ = strconv.Atoi(new_update_values[int(middle_index)])

		}
		if middle_value != -1 {
			// fmt.Println("Adding middle value", middle_value)
			total += middle_value
		}

	}

	return float64(total)
}
