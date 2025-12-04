package main

import (
	"fmt"
	"strconv"

	"github.com/harphield/advent2025/inputreader"
)

func main() {
	rows := inputreader.ReadInputFile()

	sum := 0

	for _, row := range rows {
		first := 0
		second := 0
		length := len(row)

		for i, c := range row {
			joltage, err := strconv.ParseInt(string(c), 10, 64)

			if err != nil {
				panic("parse int failed")
			}

			if i != length-1 && int(joltage) > first {
				first = int(joltage)
				second = 0
			} else if int(joltage) > second {
				second = int(joltage)
			}
		}

		fmt.Println(first, second)
		sum += first*10 + second
	}

	fmt.Println("RESULT PART 1: ", sum)
}
