package main

import (
	funcs "day5/funcs"
	"fmt"
	"utilities"
)

func main() {
	file_content1 := utilities.FileReader("data_files/part1.txt")
	answer1 := funcs.Parse_answer_one(file_content1)
	fmt.Printf("Day 5 Part 1 Answer: %f\n", answer1)

	file_content2 := utilities.FileReader("data_files/part1.txt")
	answer2 := funcs.Parse_answer_two(file_content2)
	fmt.Printf("Day 5 Part 2 Answer: %f\n", answer2)
}
