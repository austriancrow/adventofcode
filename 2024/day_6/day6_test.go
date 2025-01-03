package main

import (
	funcs "day6/funcs"
	"testing"
	"utilities"
)

func TestDay2Part1(t *testing.T) {
	test_data := utilities.FileReader("data_files/test_part1.txt")
	expected_answer := 41.0
	answer := funcs.Parse_answer_one(test_data)
	if answer != expected_answer {
		t.Fatalf("Part1 Test does not conform to expected answer: %f / %f\n", answer, expected_answer)
	}
}

func TestDay2Part2(t *testing.T) {
	test_data := utilities.FileReader("data_files/test_part1.txt")
	expected_answer := 6.0
	answer := funcs.Parse_answer_two(test_data)
	if answer != expected_answer {
		t.Fatalf("Part2 Test does not conform to expected answer: %f / %f\n", answer, expected_answer)
	}
}
