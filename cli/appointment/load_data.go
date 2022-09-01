package appointment

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Record struct {
	Name           string
	Timezone       string
	DayOfWeek      string
	AvailableAt    string
	AvailableUntil string
}

func LoadData() ([]Record, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(dir + "/data.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var record []Record
	for i, line := range data {
		if i > 0 {
			var rec Record
			for j, field := range line {
				switch j {
				case 0:
					rec.Name = field
				case 1:
					rec.Timezone = field
				case 2:
					rec.DayOfWeek = field
				case 3:
					rec.AvailableAt = field
				case 4:
					rec.AvailableUntil = field
				}
			}
			record = append(record, rec)
		}
	}
	fmt.Println(record)

	return record, nil
}
