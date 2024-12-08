package day6

import (
	"slices"
	"strings"
)

type Direction [2]int

var directions = []Direction{Direction{0, -1}, Direction{1, 0}, Direction{0, +1}, Direction{-1, 0}}

type Field struct {
	x          int
	y          int
	passable   bool
	visited    bool
	directions []int
}

type Map struct {
	fields             [][]*Field
	current_x          int
	current_y          int
	direction          int
	start_x            int
	start_y            int
	count_obstructions int
	obstructions       [][]int
}

func NewMap(data []byte) *Map {
	lines := strings.Split(string(data), "\n")
	newMap := &Map{
		count_obstructions: 0,
		direction:          0,
		obstructions:       make([][]int, 0),
	}
	newMap.fields = make([][]*Field, len(lines))
	for y, line := range lines {
		newMap.fields[y] = make([]*Field, len(line))
		for x, char := range line {
			var direction []int
			if char == '^' {
				direction = append(direction, 0)
			}
			newMap.fields[y][x] = &Field{
				x:          x,
				y:          y,
				passable:   slices.Contains([]rune{'.', '^'}, char),
				visited:    char == '^',
				directions: direction,
			}
			if char == '^' {
				newMap.current_x = x
				newMap.current_y = y
				newMap.start_x = x
				newMap.start_y = y
				newMap.direction = 0
			}
		}
	}
	return newMap
}

func (m *Map) Move() bool {
	next_x := m.current_x + directions[m.direction][0]
	next_y := m.current_y + directions[m.direction][1]
	if next_x < 0 || next_x >= len(m.fields[0]) || next_y < 0 || next_y >= len(m.fields) {
		return false
	}
	if m.fields[next_y][next_x].passable {
		m.current_x = next_x
		m.current_y = next_y
		m.fields[next_y][next_x].visited = true

	} else {
		m.direction = (m.direction + 1) % 4

	}
	return true
}

func (m *Map) Move2() bool {
	next_x := m.current_x + directions[m.direction][0]
	next_y := m.current_y + directions[m.direction][1]
	if next_x < 0 || next_x >= len(m.fields[0]) || next_y < 0 || next_y >= len(m.fields) {
		return false
	}
	if m.fields[next_y][next_x].passable {
		if next_x != m.start_x || next_y != m.start_y {
			if m.check_obstruction(m.fields[m.current_y][m.current_x], (m.direction+1)%4, 0) {
				m.count_obstructions++
				m.obstructions = append(m.obstructions, []int{next_x, next_y, m.direction})
			}
		}
		m.current_x = next_x
		m.current_y = next_y
		m.fields[next_y][next_x].visited = true
		m.fields[next_y][next_x].directions = append(m.fields[next_y][next_x].directions, (m.direction))
	} else {
		m.direction = (m.direction + 1) % 4
	}
	return true
}

func (m *Map) check_obstruction(start_field *Field, start_direction int, turns int) bool {
	if turns > 100000 {
		return false
	}
	next_x := start_field.x + directions[start_direction][0]
	next_y := start_field.y + directions[start_direction][1]
	next_direction := start_direction
	if next_x < 0 || next_x >= len(m.fields[0]) || next_y < 0 || next_y >= len(m.fields) {
		return false
	}
	if slices.Contains(m.fields[next_y][next_x].directions, start_direction) {
		return true
	}
	if !m.fields[next_y][next_x].passable {
		next_direction = (start_direction + 1) % 4
		next_x = start_field.x
		next_y = start_field.y
		turns += 1
	}
	return m.check_obstruction(m.fields[next_y][next_x], next_direction, turns)
}

func (m *Map) CountVisited() int {
	count := 0
	for _, row := range m.fields {
		for _, field := range row {
			if field.visited {
				count++
			}
		}
	}
	return count
}

func (m *Map) Print() {
	for _, row := range m.fields {
		for _, field := range row {
			var key string
			if field.visited {
				key = "X"
			} else if !field.passable {
				key = "#"
			} else {
				key = "."
			}
			for _, obstruction := range m.obstructions {
				if obstruction[0] == field.x && obstruction[1] == field.y {
					if obstruction[2] == 0 {
						key = "0"
					} else if obstruction[2] == 1 {
						key = "1"
					} else if obstruction[2] == 2 {
						key = "2"
					} else if obstruction[2] == 3 {
						key = "3"
					}
				}
			}
			print(key)
		}
		println()
	}
}

func (m *Map) count_unique_obstructions() int {
	seen := make(map[[2]int]bool)
	for _, obstruction := range m.obstructions {
		if !seen[[2]int{obstruction[0], obstruction[1]}] {
			seen[[2]int{obstruction[0], obstruction[1]}] = true
		}
	}
	return len(seen)
}
