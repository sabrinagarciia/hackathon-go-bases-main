package file

import (
	"encoding/csv"
	"io"
	"os"
	//"strings"

	"github.com/sabrinagarciia/hackathon-go-bases-main/internal/service"
)

type File struct {
	path string
}

// apunto a la dirección de la estructura (struct) File
// esto es para actualizar directamente la estructura y no una copia

func (f *File) Read() (/*[]service.Ticket*/ []string, error) {
	// abro el archivo desde su path
	file, err := os.Open("../../tickets.csv")

	if err != nil {
		return nil, err
	}

	// creo el lector del archivo
	reader := csv.NewReader(file)

	// leo linea por linea
	for {
		records, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		return records, nil
	}

	//info := strings.Split(string(file), "\n")

	// for {
	// 	records, err := reader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	for _, record := range records {
	// 		ticket := service.Ticket {
	// 			Names: record[0],
	// 		}
	// 	}
	// }
	return nil, nil
}

func (f *File) Write(service.Ticket) error {
	return nil
}

