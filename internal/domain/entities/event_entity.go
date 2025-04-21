package entities

import (
    "time"

    "github.com/google/uuid"
)

type EventModality string

const (
    Presencial EventModality = "presencial"
    Virtual    EventModality = "virtual"
    Hibrido    EventModality = "hibrido"
)

type Event struct {
    ID                       string
    Name                     string
    Description              string
    Address                  string
    MapUrl                   string
    Date                     time.Time
    Modality                 EventModality
    CancellationPolicy       string
    ParticipantEditionPolicy string
    TicketType               string
    TicketPrice              float64
    TicketQuantity           int
}

func NewEvent(
    name, description, address, mapUrl string,
    date time.Time,
    modality EventModality,
    cancellationPolicy, participantEditionPolicy, ticketType string,
    ticketPrice float64,
    ticketQuantity int,
    id string,
) *Event {
    if id == "" {
        id = uuid.New().String()
    }

    return &Event{
        ID:                       id,
        Name:                     name,
        Description:              description,
        Address:                  address,
        MapUrl:                   mapUrl,
        Date:                     date,
        Modality:                 modality,
        CancellationPolicy:       cancellationPolicy,
        ParticipantEditionPolicy: participantEditionPolicy,
        TicketType:               ticketType,
        TicketPrice:              ticketPrice,
        TicketQuantity:           ticketQuantity,
    }
}
