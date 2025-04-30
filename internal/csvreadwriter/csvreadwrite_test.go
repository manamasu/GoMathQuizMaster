package csvreadwriter

import (
	"bytes"
	"encoding/csv"
	"os"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	tmp, err := os.CreateTemp("", "test.csv")

	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())

}

func TestReadCSVRecords(t *testing.T) {
	// Create a temporary CSV file
	tmpfile, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	//csv formatting to keep it tidy and fix issue with strcon.Atoi(" ") which would lead to a read issue of a row with an empty string
	csvContent := strings.Join([]string{
		"num1,num2,solution",
		"1,2,3",
		"4,5,9",
	}, "\n")

	if _, err := tmpfile.Write([]byte(csvContent)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close() // After writing we have to close, so we can read

	records, err := ReadCSVRecords(tmpfile.Name()) //checking by using tmpFile
	if err != nil {
		t.Fatalf("ReadCSVRecords returned error: %v", err)
	}

	expected := []Record[int]{
		{Record: []int{1, 2, 3}},
		{Record: []int{4, 5, 9}},
	}

	//Making sure it matches the len of our csvContent and expected
	if len(records) != len(expected) {
		t.Fatalf("Expected %d records, got %d", len(expected), len(records))
	}

	//check records and expected for the same values, if not then test-error
	for i := range records {
		for j := range records[i].Record {
			if records[i].Record[j] != expected[i].Record[j] {
				t.Errorf("Mismatch at record %d, index %d: got %d, want %d", i, j, records[i].Record[j], expected[i].Record[j])
			}
		}
	}
}

func TestWriteCSVRecord(t *testing.T) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	rec := Record[int]{Record: []int{7, 8, 15}}
	WriteCSVRecord(writer, rec)

	output := buf.String()
	expected := "7,8,15\n"

	if output != expected {
		t.Errorf("Unexpected CSV output:\nGot:  %q\nWant: %q", output, expected)
	}
}
