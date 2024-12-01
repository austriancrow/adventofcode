package main

import (
	funcs "day1/funcs"
	"fmt"
	"time"
	"utilities"
)

func main() {
	defer utilities.TimeTrack(time.Now(), "Day 1")
	file_content1 := utilities.FileReader("data_files/part1.txt")
	answer1 := funcs.Parse_answer_one(file_content1)
	fmt.Printf("Day 1 Part 1 Answer: %f\n", answer1)
}
