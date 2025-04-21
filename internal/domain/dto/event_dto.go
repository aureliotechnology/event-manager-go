package dto

import (
    "time"

    "github.com/go-playground/validator/v10"
    "event-manager-go/internal/domain/entities"
)

var validate = validator.New()

type CreateEventDTO struct {
    Name                     string                 `json:"name" validate:"required"`
    Description              string                 `json:"description" validate:"required"`
    Address                  string                 `json:"address" validate:"required"`
    MapUrl                   string                 `json:"mapUrl" validate:"omitempty,url"`
    Date                     time.Time              `json:"date" validate:"required"`
    Modality                 entities.EventModality `json:"modality" validate:"required,oneof=presencial virtual hibrido"`
    CancellationPolicy       string                 `json:"cancellationPolicy" validate:"required"`
    ParticipantEditionPolicy string                 `json:"participantEditionPolicy" validate:"required"`
    TicketType               string                 `json:"ticketType" validate:"required"`
    TicketPrice              float64                `json:"ticketPrice" validate:"required,gt=0"`
    TicketQuantity           int                    `json:"ticketQuantity" validate:"required,gt=0"`
}

func (dto *CreateEventDTO) Validate() error {
    return validate.Struct(dto)
}

type EventDTO struct {
    ID                       string                 `json:"id"`
    Name                     string                 `json:"name"`
    Description              string                 `json:"description"`
    Address                  string                 `json:"address"`
    MapUrl                   string                 `json:"mapUrl"`
    Date                     time.Time              `json:"date"`
    Modality                 entities.EventModality `json:"modality"`
    CancellationPolicy       string                 `json:"cancellationPolicy"`
    ParticipantEditionPolicy string                 `json:"participantEditionPolicy"`
    TicketType               string                 `json:"ticketType"`
    TicketPrice              float64                `json:"ticketPrice"`
    TicketQuantity           int                    `json:"ticketQuantity"`
}

// UpdateEventDTO representa os dados de entrada para atualização de um evento.
// Todos os campos são opcionais.
type UpdateEventDTO struct {
    Name                     *string                 `json:"name,omitempty" validate:"omitempty,min=1"`
    Description              *string                 `json:"description,omitempty" validate:"omitempty,min=1"`
    Address                  *string                 `json:"address,omitempty" validate:"omitempty,min=1"`
    MapUrl                   *string                 `json:"mapUrl,omitempty" validate:"omitempty,url"`
    Date                     *time.Time              `json:"date,omitempty"`
    Modality                 *entities.EventModality `json:"modality,omitempty" validate:"omitempty,oneof=presencial virtual hibrido"`
    CancellationPolicy       *string                 `json:"cancellationPolicy,omitempty" validate:"omitempty,min=1"`
    ParticipantEditionPolicy *string                 `json:"participantEditionPolicy,omitempty" validate:"omitempty,min=1"`
    TicketType               *string                 `json:"ticketType,omitempty" validate:"omitempty,min=1"`
    TicketPrice              *float64                `json:"ticketPrice,omitempty" validate:"omitempty,gt=0"`
    TicketQuantity           *int                    `json:"ticketQuantity,omitempty" validate:"omitempty,gt=0"`
}

// Validate executa a validação dos campos do UpdateEventDTO.
func (dto *UpdateEventDTO) Validate() error {
    return validate.Struct(dto)
}
