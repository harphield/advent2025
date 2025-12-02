package inputreader

import (
	"bufio"
	"os"
	"path/filepath"
)

func ReadInputFile() []string {
	path := filepath.Join(".", "input.txt")

	f, err := os.Open(path)

	check(err)

	scanner := bufio.NewScanner(f)

	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	f.Close()

	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
