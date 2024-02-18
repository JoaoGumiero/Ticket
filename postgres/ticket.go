package postgres

import "time"

type Ticket struct {
	Id int
	Number int
	EntryDate time.Time
	ResolutionDate time.Time
	Country string
	Type string
	Category string
	Status string
	ResolutionDetails string
	OriginalQueue string
	DestinationQueue string
}