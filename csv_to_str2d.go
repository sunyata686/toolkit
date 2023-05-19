package toolkit

import (
	"encoding/csv"
	"os"
)

func CsvToStr2d(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return [][]string{}, err
	}

	c := csv.NewReader(f)
	c.LazyQuotes = true
	return c.ReadAll()
}
