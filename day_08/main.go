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
	z int
}

type distance struct {
	id_1     int
	id_2     int
	distance float64
}

func main() {
	rows := inputreader.ReadInputFile()

	junctions := make([]coords, 0)

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

		z, err := strconv.ParseInt(split[2], 10, 64)

		if err != nil {
			panic("parseint failed")
		}

		junctions = append(junctions, coords{int(x), int(y), int(z)})
	}

	// fmt.Println(junctions)

	distances := make([]distance, 0)

	for id_1, c := range junctions {
		for id_2 := id_1 + 1; id_2 < len(junctions); id_2++ {
			c2 := junctions[id_2]

			d := distance{
				id_1,
				id_2,
				math.Sqrt(
					math.Pow(float64(c.x)-float64(c2.x), 2) +
						math.Pow(float64(c.y)-float64(c2.y), 2) +
						math.Pow(float64(c.z)-float64(c2.z), 2),
				),
			}

			distances = append(distances, d)
		}
	}

	slices.SortStableFunc(distances, func(a, b distance) int {
		if a.distance < b.distance {
			return -1
		}

		if b.distance > a.distance {
			return 1
		}

		return 0
	})

	connections := make([][]bool, 0)
	for i := 0; i < len(junctions); i++ {
		connections = append(connections, make([]bool, len(junctions)))
	}

	for i := range 1000 {
		d := distances[i]

		if !connections[d.id_1][d.id_2] {
			connections[d.id_1][d.id_2] = true
			connections[d.id_2][d.id_1] = true

			// fmt.Println(d.id_1, " ", d.id_2)
		}
	}

	circuits := make([][]int, len(junctions)) // in the beginning, all of them are separate circuits
	for i, _ := range junctions {
		circuits[i] = append(circuits[i], i)
	}

	for id_1, c := range connections {
		for id_2, v := range c {
			if v {
				// merge circuits where id_1 and id_2 are located
			CIRCUIT_FINDER:
				for i := range circuits {
					if slices.Contains(circuits[i], id_1) && !slices.Contains(circuits[i], id_2) {
						for j := range circuits {
							if slices.Contains(circuits[j], id_2) {
								circuits[i] = append(circuits[i], circuits[j]...)
								circuits[j] = make([]int, 0)
								break CIRCUIT_FINDER
							}
						}
					}
				}
			}
		}
	}

	slices.SortStableFunc(circuits, func(a, b []int) int {
		if len(a) > len(b) {
			return -1
		}

		if len(a) < len(b) {
			return 1
		}

		return 0
	})

	// fmt.Println(circuits)

	result := 1
	for _, c := range circuits[:3] {
		result *= len(c)
	}

	fmt.Println("RESULT PART 1: ", result)
}
