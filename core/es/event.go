package es

import (
	"time"
)

// Event is the structure for all our events
type Event struct {
	AggregateID        string
	Typology  Typology
	Payload   interface{}
	CreatedAt time.Time
	Index     uint
}

// Typology of an event
type Typology string

// typology types
const (
	Create Typology = "create"
	Put    Typology = "put"
	Delete Typology = "delete"
)
