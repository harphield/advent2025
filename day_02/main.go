package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/harphield/advent2025/inputreader"
)

func main() {
	rows := inputreader.ReadInputFile()

	ranges := strings.SplitSeq(rows[0], ",")

	sum := 0
	sum2 := 0

	for r := range ranges {
		limits := strings.Split(r, "-")
		start, err := strconv.ParseInt(limits[0], 10, 64)

		if err != nil {
			panic("parseint failed")
		}

		end, err := strconv.ParseInt(limits[1], 10, 64)

		if err != nil {
			panic("parseint failed")
		}

		for i := start; i <= end; i++ {
			num_str := strconv.Itoa(int(i))

			length := len(num_str)

			// part 2: repeating pattern search
		MAIN:
			for pattern_length := 1; pattern_length <= length/2; pattern_length++ {
				pattern := num_str[0:pattern_length]

				next := pattern_length

				for {
					if num_str[next:next+pattern_length] != pattern {
						continue MAIN
					}

					next += pattern_length
					if next == length {
						break
					}

					if next > length || next+pattern_length > length {
						continue MAIN
					}
				}

				// found an invalid ID
				sum2 += int(i)
				break
			}

			// we don't care about strings with an odd length for part 1
			if length%2 != 0 {
				continue
			}

			if num_str[0:length/2] == num_str[length/2:length] {
				sum += int(i)
			}
		}
	}

	fmt.Println("RESULT PART 1: ", sum)
	fmt.Println("RESULT PART 2: ", sum2)
}
