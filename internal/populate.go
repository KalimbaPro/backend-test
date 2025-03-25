package internal

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	charmLog "github.com/charmbracelet/log"
)

func parseCSV(filename string) ([]Breed) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var allBreeds [][]string
	for {
		breed, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		allBreeds = append(allBreeds, breed)
	}

	var breeds []Breed
	for _, line := range allBreeds[1:] { // Start at line 2 to avoid the fields specifications
		currentId, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Println("Can't convert this to an int!")
		}
		current_averageMaleAdultWeight, err := strconv.Atoi(line[4])
		if err != nil {
			fmt.Println("Can't convert this to an int!")
		}
		current_averageFemaleAdultWeight, err := strconv.Atoi(line[5])
		if err != nil {
			fmt.Println("Can't convert this to an int!")
		}
		breed := Breed{
			Id: currentId,
			Species: line[1],
			PetSize: line[2],
			Name: line[3],
			AverageMaleAdultWeight: current_averageMaleAdultWeight,
			AverageFemaleAdultWeight: current_averageFemaleAdultWeight,
		}
		breeds = append(breeds, breed)
	}
	return breeds
}

func PopulateDatabase(filename string, db *sql.DB) error {
	breeds := parseCSV(filename)

	logger := charmLog.NewWithOptions(os.Stderr, charmLog.Options{
		Formatter:       charmLog.TextFormatter,
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "🧑‍💻 backend-test",
		Level:           charmLog.DebugLevel,
	})

	var isError bool
	var err error
	for index, breed := range breeds {
		query := "INSERT INTO breeds(id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight) VALUES(?, ?, ?, ?, ?, ?)"

		_, err = db.Exec(query, breed.Id, breed.Species, breed.PetSize, breed.Name, breed.AverageMaleAdultWeight, breed.AverageFemaleAdultWeight)
		if err != nil {
			isError = true
			logger.Errorf("Populate error at: %d", index)
		}
	}
	if isError {
		return err
	}
	return nil
}
