package main

import (
	"github.com/sabrinagarciia/hackathon-go-bases-main/internal/service"
	//"github.com/sabrinagarciia/hackathon-go-bases-main/internal/file"
	// "encoding/csv"
	// "fmt"
	// "io"
	// "os"
)

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)

	// tfile, err := os.Open("tickets.csv")

	// if err != nil {
	// 	panic("El path es incorrecto")
	// }

	// reader := csv.NewReader(tfile)
	// for {
	// 	records, err := reader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	if err != nil {
	// 		panic("error")
	// 	}

		// for value := range record {
		// 	fmt.Printf("%s\n", record[value])
		// }
		// for _, record := range records {
		// 	ticket := service.Ticket {
		// 		Names: record[1],
		// 		Email: record[2],
		// 		Destination: record[3],
		// 	}
		// 	fmt.Printf("%s %s %s \n", ticket.Names, ticket.Email, ticket.Destination)
		// }
	// }
}
