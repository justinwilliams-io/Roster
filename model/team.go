package model

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/dnlo/struct2csv"
	"github.com/google/uuid"
)

type Team struct {
    Name   string
	Roster []Player
}

func (t *Team) CreateCsv() string {
    id := uuid.New()
	fileName := t.Name + "-" + id.String() + ".csv"

	_, err := os.Stat("./files")
	if os.IsNotExist(err) {
		os.MkdirAll("./files", os.ModePerm)
	}

	file, err := os.Create("./files/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	encoder := struct2csv.New()
	rows, err := encoder.Marshal(t.Roster)
	if err != nil {
		log.Fatal(err)
	}

	writer.WriteAll(rows)

	return fileName
}
