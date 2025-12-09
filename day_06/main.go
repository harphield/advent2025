package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/harphield/advent2025/inputreader"
)

func main() {
	re_numbers := regexp.MustCompile(`([0-9]+)`)
	re_symbols := regexp.MustCompile(`([\+\*])+`)

	rows := inputreader.ReadInputFile()

	var columns [][]int
	var part2_columns [][]string
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

			if len(part2_columns) == 0 {
				for range len(row) {
					part2_columns = append(part2_columns, make([]string, 0))
				}
			}

			for i, n := range number_matches {
				number, err := strconv.ParseInt(n, 10, 64)

				if err != nil {
					panic("parseint failed")
				}

				columns[i] = append(columns[i], int(number))
			}

			// in part 2 I fill all characters in the row
			for i2, s := range row {
				if i2 >= len(part2_columns) {
					part2_columns = append(part2_columns, make([]string, 0))
				}
				part2_columns[i2] = append(part2_columns[i2], string(s))
			}
		} else {
			// this only appears once in the end
			operators = re_symbols.FindAllString(row, -1)
		}
	}

	// part 1
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

	// part 2
	current_operator_index := len(operators) - 1
	column_index := len(part2_columns) - 1
	sum2 := 0
	multi_value := 1
	for {
		number := build_number_from_column(part2_columns[column_index])

		if number == 0 {
			if current_operator_index-1 < 0 {
				break
			}

			if operators[current_operator_index] == "*" {
				sum2 += multi_value
				multi_value = 1
			}

			current_operator_index--
			column_index--
			continue
		}

		if operators[current_operator_index] == "+" {
			sum2 += number
		} else {
			multi_value *= number
		}

		if column_index-1 < 0 {
			if operators[current_operator_index] == "*" {
				sum2 += multi_value
			}
			break
		}

		column_index--
	}

	fmt.Println("RESULT PART 1: ", sum)
	fmt.Println("RESULT PART 2: ", sum2)
}

func build_number_from_column(column []string) int {
	result := 0
	decimals := 0
	for i := len(column) - 1; i >= 0; i-- {
		number, err := strconv.ParseInt(column[i], 10, 64)

		if err != nil {
			continue
		}

		result += int(number * int64(math.Pow10(decimals)))
		decimals++
	}

	return result
}
