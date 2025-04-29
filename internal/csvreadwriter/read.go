package csvreadwriter

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func CreateCSVReader(filename string) (*csv.Reader, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	return r, nil
}

func ReadCSVRecords(filename string) ([]Record[int], error) {

	r, err := CreateCSVReader(filename)

	if err != nil {
		fmt.Println("Something went wrong by creating CSVReader")
	}
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
