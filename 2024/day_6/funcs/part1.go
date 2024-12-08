package day6

import (
	"fmt"
	"time"
	"utilities"
)

func Parse_answer_one(file_content []byte) float64 {
	defer utilities.TimeTrack(time.Now(), "Day 6 Part 1")
	total := 0
	moveMap := NewMap(file_content)
	for moveMap.Move() {
		total++
	}
	fmt.Println("Total steps: ", total)
	visited_spaces := moveMap.CountVisited()
	moveMap.Print()
	return float64(visited_spaces)
}
