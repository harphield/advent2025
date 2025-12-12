package main

import (
	"fmt"

	"github.com/harphield/advent2025/inputreader"
)

type coords struct {
	x int
	y int
}

const EMPTY = 0
const START = 1
const SPLITTER = 2
const BEAM = 3

func main() {
	rows := inputreader.ReadInputFile()

	x := 0
	y := 0

	var hitmap [][]int

	width := 0
	splitters := 0
	start_x := 0
	start_y := 0

	for _, row := range rows {
		if width == 0 {
			width = len(row)
		}

		hitmap = append(hitmap, make([]int, 0))
		for _, c := range row {
			switch c {
			case 'S':
				// start
				start_x = x
				start_y = y
				hitmap[y] = append(hitmap[y], START)
			case '^':
				// splitter
				hitmap[y] = append(hitmap[y], SPLITTER)
				// check if a ray is coming in from above
				if hitmap[y-1][x] == BEAM {
					if x > 0 {
						hitmap[y][x-1] = BEAM
					}

					splitters++
				}
			default:
				if x > 0 && hitmap[y][x-1] == SPLITTER && hitmap[y-1][x] != BEAM {
					hitmap[y] = append(hitmap[y], BEAM)
				} else if y > 0 && (hitmap[y-1][x] == BEAM || hitmap[y-1][x] == START) {
					hitmap[y] = append(hitmap[y], BEAM)
				} else {
					hitmap[y] = append(hitmap[y], EMPTY)
				}
			}

			x++
		}

		y++
		x = 0
	}

	// reset BEAMs from the hitmap
	for i := 0; i < len(hitmap); i++ {
		for j := 0; j < len(hitmap[i]); j++ {
			if hitmap[i][j] == BEAM {
				hitmap[i][j] = EMPTY
			}
		}
	}

	for _, row := range hitmap {
		fmt.Println(row)
	}

	counted := make(map[coords]int)
	timelines := shoot_lazor(hitmap, start_x, start_y+1, counted)

	fmt.Println("RESULT PART 1: ", splitters)
	fmt.Println("RESULT PART 2: ", timelines)
}

func shoot_lazor(hitmap [][]int, x int, y int, counted map[coords]int) int {
	if counted[coords{x, y}] > 0 {
		return counted[coords{x, y}]
	}

	if y == len(hitmap) {
		return 1
	}

	result := 0

	if hitmap[y][x] != SPLITTER {
		hitmap[y][x] = BEAM
		result += shoot_lazor(hitmap, x, y+1, counted)
		hitmap[y][x] = EMPTY
	} else {
		if x+1 < len(hitmap[y]) && hitmap[y-1][x+1] != BEAM {
			hitmap[y][x+1] = BEAM
			result += shoot_lazor(hitmap, x+1, y, counted)
			hitmap[y][x+1] = EMPTY
		}

		if x > 0 && hitmap[y-1][x-1] != BEAM {
			hitmap[y][x-1] = BEAM
			result += shoot_lazor(hitmap, x-1, y, counted)
			hitmap[y][x-1] = EMPTY
		}
	}

	// saving this result so I don't need to repeat this path :)
	counted[coords{x, y}] = result

	return result
}
