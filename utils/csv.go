package utils

import (
	"encoding/csv"
	"io"
	"os"
)

// LoadCSV CSVを読み込む
func LoadCSV(path string) [][]string {
	bulkCount := 100

	file, _ := os.Open(path)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	lines := make([][]string, 0, bulkCount)
	for {
		isLast := false
		for i := 0; i < bulkCount; i++ {
			line, err := reader.Read()
			if err == io.EOF {
				isLast = true
				break
			} else if err != nil {
				panic(err)
			}
			lines = append(lines, line)
		}

		if isLast {
			break
		}
	}
	return lines
}
