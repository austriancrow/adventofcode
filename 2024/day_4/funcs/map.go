package day4

import "strings"

type Map struct {
	fields [][]Field
}

type Field struct {
	value     string
	neighbors [8]*Field
}

func NewMap(file_content []byte) *Map {
	m := &Map{}
	lines := strings.Split(string(file_content), "\n")
	m.fields = make([][]Field, len(lines))
	for y, line := range strings.Split(string(file_content), "\n") {
		m.fields[y] = make([]Field, len(line))
		for x, letter := range line {
			// Create a new field
			m.fields[y][x] = Field{value: string(letter)}
		}
	}

	for y, row := range m.fields {
		for x, field := range row {
			/*
				0 north
				1 north-east
				2 east
				3 south-east
				4 south
				5 south-west
				6 west
				7 north-west
			*/
			if y > 0 {
				field.neighbors[0] = &m.fields[y-1][x]
				if x > 0 {
					field.neighbors[7] = &m.fields[y-1][x-1]
				}
				if x < len(row)-1 {
					field.neighbors[1] = &m.fields[y-1][x+1]
				}
			}
			if x > 0 {
				field.neighbors[6] = &m.fields[y][x-1]
			}
			if x < len(row)-1 {
				field.neighbors[2] = &m.fields[y][x+1]
			}
			if y < len(m.fields)-1 {
				field.neighbors[4] = &m.fields[y+1][x]
				if x > 0 {
					field.neighbors[5] = &m.fields[y+1][x-1]
				}
				if x < len(row)-1 {
					field.neighbors[3] = &m.fields[y+1][x+1]
				}
			}
		}
	}

	return m
}
