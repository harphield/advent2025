package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/harphield/advent2025/inputreader"
)

func main() {
	rows := inputreader.ReadInputFile()

	reading_ids := false

	var ranges [][]uint64
	var ids []uint64

	for _, row := range rows {
		if row == "" {
			reading_ids = true
			continue
		}

		if !reading_ids {
			// read ranges
			split := strings.Split(row, "-")

			start, err := strconv.ParseInt(split[0], 10, 64)

			if err != nil {
				panic("parse int failed")
			}

			end, err := strconv.ParseInt(split[1], 10, 64)

			if err != nil {
				panic("parse int failed")
			}

			ranges = append(ranges, []uint64{uint64(start), uint64(end)})

		} else {
			// read IDs
			id, err := strconv.ParseInt(row, 10, 64)

			if err != nil {
				panic("parse int failed")
			}

			ids = append(ids, uint64(id))
		}

	}

	// sort the ranges
	slices.SortFunc(ranges, func(a []uint64, b []uint64) int {
		if a[0] == b[0] && a[1] == b[1] {
			return 0
		}

		if a[0] < b[0] || (a[0] == b[0] && a[1] < b[1]) {
			return -1
		}

		if a[0] > b[0] || (a[0] == b[0] && a[1] > b[1]) {
			return 1
		}

		return 0
	})

	var optimized_ranges [][]uint64

	ranges_count := len(ranges)
	i := 0
	new_range := ranges[0]
	for {
		if i+1 >= ranges_count {
			optimized_ranges = append(optimized_ranges, new_range)
			break
		}

		if ranges[i][1] >= ranges[i+1][0] {
			if ranges[i+1][1] > new_range[1] {
				new_range[1] = ranges[i+1][1]
			}
		} else {
			optimized_ranges = append(optimized_ranges, new_range)
			new_range = ranges[i+1]
		}

		i++
	}

	fresh := 0

IDLOOP:
	for _, id := range ids {
		for _, r := range optimized_ranges {
			if id >= r[0] && id <= r[1] {
				fresh++
				continue IDLOOP
			}
		}
	}

	fmt.Println("RESULT PART 1: ", fresh)

	fresh_count := uint64(0)
	previous := []uint64{0, 0}
	for _, r := range optimized_ranges {
		if previous[1] == r[0] {
			// for some reason I did not join some ranges correctly, so this fixes it D:
			fresh_count--
		}

		fresh_count += r[1] - r[0] + 1
		previous = r
	}

	fmt.Println("RESULT PART 2: ", fresh_count)
}
