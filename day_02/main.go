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

	for r := range ranges {
		fmt.Println(r)

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

			// we don't care about strings with an odd length
			if length%2 != 0 {
				continue
			}

			if num_str[0:length/2] == num_str[length/2:length] {
				sum += int(i)
			}
		}
	}

	fmt.Println("RESULT: ", sum)
}
