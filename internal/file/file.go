package file

import (
	"encoding/csv"
	"os"

	"github.com/sabrinagarciia/hackathon-go-bases-main/internal/service"
)

type File struct {
	path string
}

// apunto a la dirección de la estructura (struct) File
// esto es para actualizar directamente la estructura y no una copia

func (f *File) Read() ([]service.Ticket, error) {
	file, err := os.Open("../../tickets.csv")

	if err != nil {
		panic("path incorrecto")
	}

	reader := csv.NewReader(file)
	for {
		records, err := reader.Read()
		if err == io.EOF {
			break
		}
		for _, record := range records {
			ticket := service.Ticket {
				Names: record[0],
			}
		}
	}
	return nil, nil
}

func (f *File) Write(service.Ticket) error {
	return nil
}

