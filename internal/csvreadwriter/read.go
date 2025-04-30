package csvreadwriter

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// CreateCSVReader opens the specified file and returns a CSV reader and the file handle.
// The caller is responsible for closing the returned file when done.
func CreateCSVReader(filename string) (*csv.Reader, *os.File, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}

	r := csv.NewReader(f)
	return r, f, nil
}

// ReadCSVRecords reads a CSV file with a header and returns a slice of Record[int].
// Assumes each row contains only integer values after the header.
func ReadCSVRecords(filename string) ([]Record[int], error) {

	r, f, err := CreateCSVReader(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSV reader: %w", err)
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
		if len(row) == 0 {
			continue // skipping blank lines
		}

		var nums []int

		for _, field := range row {
			field = strings.TrimSpace(field)
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
