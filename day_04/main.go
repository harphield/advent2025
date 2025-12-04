package main

import (
	"fmt"
	"math"

	"github.com/harphield/advent2025/inputreader"
)

func main() {
	rows := inputreader.ReadInputFile()

	var grid []bool
	width := 0
	height := 0

	for _, row := range rows {
		if width == 0 {
			width = len(row)
		}

		for _, c := range row {
			if c == '.' {
				grid = append(grid, false)
			} else {
				grid = append(grid, true)
			}
		}

		height++
	}

	draw_grid(grid, width)

	can_be_picked_up := 0
	loops := 0
	var to_remove []int

	for {
		for i, v := range grid {
			if !v {
				continue
			}

			x := i % width
			y := int(math.Floor(float64(i / width)))

			rolls_around := 0

			// left
			if x > 0 && grid[coords_to_index(x-1, y, width)] == true {
				rolls_around++
			}
			// right
			if x < width-1 && grid[coords_to_index(x+1, y, width)] == true {
				rolls_around++
			}
			// up
			if y > 0 && grid[coords_to_index(x, y-1, width)] == true {
				rolls_around++
			}
			// down
			if y < height-1 && grid[coords_to_index(x, y+1, width)] == true {
				rolls_around++
			}
			// up left
			if x > 0 && y > 0 && grid[coords_to_index(x-1, y-1, width)] == true {
				rolls_around++
			}
			// up right
			if x < width-1 && y > 0 && grid[coords_to_index(x+1, y-1, width)] == true {
				rolls_around++
			}
			// down left
			if x > 0 && y < height-1 && grid[coords_to_index(x-1, y+1, width)] == true {
				rolls_around++
			}
			// down right
			if x < width-1 && y < height-1 && grid[coords_to_index(x+1, y+1, width)] == true {
				rolls_around++
			}

			if rolls_around < 4 {
				//fmt.Println(x, y, " can be picked up")
				can_be_picked_up++
				to_remove = append(to_remove, i)
			}
		}

		if loops == 0 {
			fmt.Println("RESULT PART 1: ", can_be_picked_up)
		}

		loops++

		if len(to_remove) > 0 {
			for _, index := range to_remove {
				grid[index] = false
			}

			to_remove = make([]int, 0)
			continue
		}

		break
	}

	fmt.Println("RESULT PART 2: ", can_be_picked_up)
}

func coords_to_index(x int, y int, width int) int {
	return (y * width) + x
}

func draw_grid(grid []bool, width int) {
	for i, v := range grid {
		if i%width == 0 {
			fmt.Println()
		}

		if v {
			fmt.Print("@")
		} else {
			fmt.Print(".")
		}
	}

	fmt.Println()
}
