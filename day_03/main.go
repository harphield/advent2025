package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/harphield/advent2025/inputreader"
)

func main() {
	rows := inputreader.ReadInputFile()

	sum := 0
	sum2 := 0

	for _, row := range rows {
		sum += solve(row, 2)
		sum2 += solve(row, 12)
	}

	fmt.Println("RESULT PART 1: ", sum)
	fmt.Println("RESULT PART 2: ", sum2)
}

func solve(row string, bat_length int) int {
	values := make([]int, bat_length)
	length := len(row)

	// we can search for values from this index
	search_from := 0

	for i, c := range row {
		joltage, err := strconv.ParseInt(string(c), 10, 64)

		if err != nil {
			panic("parse int failed")
		}

		if length-bat_length-i < 0 {
			search_from = -(length - bat_length - i)
		}

		reset_mode := false
		for vi := search_from; vi < len(values); vi++ {
			if reset_mode {
				values[vi] = 0
				continue
			}

			if int(joltage) > values[vi] {
				values[vi] = int(joltage)
				reset_mode = true
			}
		}
	}

	result := 0

	for i, value := range values {
		result += value * int(math.Pow10(bat_length-i-1))
	}

	return result
}
