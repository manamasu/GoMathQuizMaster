package csvreadwriter

import (
	"encoding/csv"
	"fmt"
	"os"
)

// CreateCSVWriter creates a file f in the current folder and passes the csv writer w through
func CreateCSVWriter(filename string) (*csv.Writer, *os.File, error) {

	f, err := os.Create(filename)

	if err != nil {
		fmt.Println("There has been an error with Creating a file. Error: ", err)
		return nil, nil, err
	}
	w := csv.NewWriter(f)

	return w, f, nil
}

// WriteCSVRecord writes a Record[T] to the given CSV writer.
// converts each field to a string before writing, and flushes (writes to csv).
// T can be any type that can be stringified using fmt.Sprint.
func WriteCSVRecord[T any](writer *csv.Writer, rec Record[T]) {
	var stringRecord []string
	for _, v := range rec.Record {
		stringRecord = append(stringRecord, fmt.Sprint(v))
	}

	err := writer.Write(stringRecord)
	if err != nil {
		fmt.Println("Error writing record to csv: ", err)
	}
	writer.Flush()
}
