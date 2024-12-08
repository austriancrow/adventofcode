package day6

import (
	"fmt"
	"time"
	"utilities"
)

func Parse_answer_two(file_content []byte) float64 {
	defer utilities.TimeTrack(time.Now(), "Day 6 Part 2")
	total := 0
	moveMap := NewMap(file_content)
	for moveMap.Move2() {
		total++
	}
	fmt.Println("Total steps: ", total)
	count_obstructions := moveMap.count_unique_obstructions()
	moveMap.Print()
	fmt.Println("-------------------")
	fmt.Println("Obstructions: ", moveMap.obstructions)
	return float64(count_obstructions)
}
