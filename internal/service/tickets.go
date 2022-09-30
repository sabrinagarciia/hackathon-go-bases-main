package service

import "errors"

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	for _, value := range b.Tickets {
		if value.Id == t.Id {
			return Ticket{}, errors.New("errorrrrr")
		}
	}

	b.Tickets = append(b.Tickets, t)
	return Ticket{}, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, value := range b.Tickets {
		if value.Id != id {
			return value, errors.New("no existe")
		}
	}

	return Ticket{}, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	if t.Id != id {
		return Ticket{}, errors.New("no se ingreso el id correspondiente al ticket")
	}

	for i, value := range b.Tickets {
		if value.Id == id {
			b.Tickets[i] = t
			return t, nil
		}
	}

	return Ticket{}, errors.New("no existe")
}

func (b *bookings) Delete(id int) (int, error) {
	return 0, nil
	//b.Tickets = []Ticket{}
}
