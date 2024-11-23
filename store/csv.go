package store

import (
	"encoding/csv"
	"errors"
	"os"
)

type CSVDatabase struct {
	data map[string]*Location
}

func InitCSVDatabase(path string) (*CSVDatabase, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	data := make(map[string]*Location)

	for _, record := range records {
		if len(record) != 3 {
			continue
		}
		ip := record[0]
		city := record[1]
		country := record[2]
		data[ip] = &Location{Country: country, City: city}

	}

	return &CSVDatabase{data: data}, nil

}

func (db *CSVDatabase) Find(ip string) (*Location, error) {
	location, exists := db.data[ip]

	if !exists {
		return nil, errors.New("IP not found")
	}

	return location, nil
}
