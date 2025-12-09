package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/harphield/advent2025/inputreader"
)

func main() {
	re_numbers := regexp.MustCompile(`([0-9]+)`)
	re_symbols := regexp.MustCompile(`([\+\*])+`)

	rows := inputreader.ReadInputFile()

	var columns [][]int
	var operators []string

	for _, row := range rows {
		number_matches := re_numbers.FindAllString(row, -1)

		if len(number_matches) > 0 {
			if len(columns) == 0 {
				// prefill with arrays
				for range number_matches {
					columns = append(columns, make([]int, 0))
				}
			}

			for i, n := range number_matches {
				number, err := strconv.ParseInt(n, 10, 64)

				if err != nil {
					panic("parseint failed")
				}

				columns[i] = append(columns[i], int(number))
			}
		} else {
			// this only appears once in the end
			operators = re_symbols.FindAllString(row, -1)
		}
	}

	fmt.Println(operators)
	fmt.Println(columns)

	sum := 0
	for i, o := range operators {
		if o == "+" {
			for _, num := range columns[i] {
				sum += num
			}
		} else {
			result := 1
			for _, num := range columns[i] {
				result *= num
			}

			sum += result
		}
	}

	fmt.Println("RESULT PART 1: ", sum)
}
