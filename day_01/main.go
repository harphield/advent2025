package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/harphield/advent2025/inputreader"
)

func main() {
	rows := inputreader.ReadInputFile()

	println(len(rows))

	re := regexp.MustCompile(`(L|R)(.+)`)

	arrow := 50
	zeroes := 0
	zeroes2 := 0

	for _, value := range rows {
		matches := re.FindAllStringSubmatch(value, -1)

		direction := matches[0][1]
		distance, err := strconv.ParseInt(matches[0][2], 10, 64)

		if err != nil {
			panic("parseint failed")
		}

		fmt.Println(direction, distance)

		div := float64(0)

		switch direction {
		case "L":
			div = float64(int(distance)+(100-arrow)) / 100

			if arrow == 0 {
				div -= 1
			}

			arrow -= int(distance)
		case "R":
			div = float64(int(distance)+arrow) / 100

			arrow += int(distance)
		}

		// part 2
		fmt.Println(div)
		zeroes2 += int(math.Floor(div))

		if arrow < 0 {
			arrow = 100 + (arrow % 100)
			if arrow == 100 {
				arrow = 0
			}
		} else if arrow > 99 {
			arrow = arrow % 100
		}

		fmt.Println(arrow)

		if arrow == 0 {
			zeroes++
		}
	}

	fmt.Println("RESULT PART 1: ", zeroes)
	fmt.Println("RESULT PART 2: ", zeroes2)
}
