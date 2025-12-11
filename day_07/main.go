package main

import (
	"fmt"

	"github.com/harphield/advent2025/inputreader"
)

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

	for _, row := range rows {
		if width == 0 {
			width = len(row)
		}

		hitmap = append(hitmap, make([]int, 0))

		for _, c := range row {
			switch c {
			case 'S':
				// start
				hitmap[y] = append(hitmap[y], START)
			case '^':
				// splitter
				hitmap[y] = append(hitmap[y], SPLITTER)
				// check if a ray is coming in from above
				if hitmap[y-1][x] == BEAM {
					if x > 0 && hitmap[y][x-1] == EMPTY {
						hitmap[y][x-1] = BEAM
					}

					splitters++
				}
			default:
				if y > 0 && (hitmap[y-1][x] == BEAM || hitmap[y-1][x] == START) {
					hitmap[y] = append(hitmap[y], BEAM)
				} else if x > 0 && hitmap[y][x-1] == SPLITTER {
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

	for _, row := range hitmap {
		fmt.Println(row)
	}

	fmt.Println("RESULT PART 1: ", splitters)
}
