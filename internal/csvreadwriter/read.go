package csvreadwriter

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func CreateCSVReader(filename string) (*csv.Reader, *os.File, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}

	r := csv.NewReader(f)
	return r, f, nil
}

func ReadCSVRecords(filename string) ([]Record[int], error) {

	r, f, err := CreateCSVReader(filename)
	if err != nil {
		fmt.Println("Something went wrong by creating CSVReader")
	}
	defer f.Close()

	// Skip header
	if _, err := r.Read(); err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	var records []Record[int]

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error reading row: %w", err)
		}

		var nums []int

		for _, field := range row {
			n, err := strconv.Atoi(field)

			if err != nil {
				return nil, fmt.Errorf("non-integer value found: %s", field)
			}
			nums = append(nums, n)
		}

		records = append(records, Record[int]{Record: nums})
	}

	return records, nil
}
