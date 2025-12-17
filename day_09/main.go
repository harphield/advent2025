package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/harphield/advent2025/inputreader"
)

type coords struct {
	x int
	y int
}

type distance struct {
	id_1     int
	id_2     int
	distance int
}

func main() {
	rows := inputreader.ReadInputFile()

	positions := make([]coords, 0)

	for _, row := range rows {
		split := strings.Split(row, ",")

		x, err := strconv.ParseInt(split[0], 10, 64)

		if err != nil {
			panic("parseint failed")
		}

		y, err := strconv.ParseInt(split[1], 10, 64)

		if err != nil {
			panic("parseint failed")
		}

		positions = append(positions, coords{
			x: int(x),
			y: int(y),
		})
	}

	// fmt.Println(positions)

	distances := make([]distance, 0)

	for id_1, c := range positions {
		for id_2 := id_1 + 1; id_2 < len(positions); id_2++ {
			c2 := positions[id_2]

			d := distance{
				id_1,
				id_2,
				find_distance(c, c2),
			}

			distances = append(distances, d)
		}
	}

	slices.SortStableFunc(distances, func(a, b distance) int {
		if a.distance > b.distance {
			return -1
		}

		if b.distance < a.distance {
			return 1
		}

		return 0
	})

	// fmt.Println(distances)

	fmt.Println("RESULT PART 1: ", find_area(positions[distances[0].id_1], positions[distances[0].id_2]))
}

func find_distance(c, c2 coords) int {
	return int(math.Abs(float64(c.x)-float64(c2.x)) + math.Abs(float64(c.y)-float64(c2.y)))
}

func find_area(c, c2 coords) int {
	return int((math.Abs(float64(c.x)-float64(c2.x)) + 1) * (math.Abs(float64(c.y)-float64(c2.y)) + 1))
}
