package toolkit

import (
	"fmt"
	"testing"
)

func TestCsvToStr2d(t *testing.T) {
	path := "test_files/testReadCsv.csv"
	str2d, err := CsvToStr2d(path)
	if err != nil {
		t.Fatal("Error while TestCsvToStr2d: ", err)
	}
	fmt.Println(str2d)
}
