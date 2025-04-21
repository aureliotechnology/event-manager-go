package dto

import (
	"testing"
	"time"

	"event-manager-go/internal/domain/entities"
)

func TestCreateEventDTO_Validate(t *testing.T) {
	// Caso válido
	validDTO := CreateEventDTO{
		Name:                     "Evento Teste",
		Description:              "Descrição de teste",
		Address:                  "Endereço de teste",
		MapUrl:                   "http://example.com",
		Date:                     time.Now().Add(24 * time.Hour),
		Modality:                 entities.EventModality("presencial"),
		CancellationPolicy:       "Política",
		ParticipantEditionPolicy: "Edição",
		TicketType:               "VIP",
		TicketPrice:              100.0,
		TicketQuantity:           10,
	}

	if err := validDTO.Validate(); err != nil {
		t.Errorf("Validação esperada com sucesso, mas ocorreu erro: %v", err)
	}

	// Caso inválido: Nome vazio
	invalidDTO := validDTO
	invalidDTO.Name = ""
	if err := invalidDTO.Validate(); err == nil {
		t.Errorf("Era esperado erro de validação para Nome vazio")
	}

	// Caso inválido: URL inválida (caso MapUrl seja fornecido)
	invalidDTO = validDTO
	invalidDTO.MapUrl = "invalid_url"
	if err := invalidDTO.Validate(); err == nil {
		t.Errorf("Era esperado erro de validação para MapUrl inválido")
	}

	// Caso inválido: TicketPrice <= 0
	invalidDTO = validDTO
	invalidDTO.TicketPrice = 0
	if err := invalidDTO.Validate(); err == nil {
		t.Errorf("Era esperado erro de validação para TicketPrice <= 0")
	}

	// Caso inválido: TicketQuantity <= 0
	invalidDTO = validDTO
	invalidDTO.TicketQuantity = 0
	if err := invalidDTO.Validate(); err == nil {
		t.Errorf("Era esperado erro de validação para TicketQuantity <= 0")
	}
}

func TestUpdateEventDTO_Validate(t *testing.T) {
	// Como todos os campos são opcionais, um UpdateEventDTO vazio deve ser válido.
	emptyDTO := UpdateEventDTO{}
	if err := emptyDTO.Validate(); err != nil {
		t.Errorf("Era esperado que um UpdateEventDTO vazio fosse válido, mas ocorreu erro: %v", err)
	}

	// Cria valores válidos para os campos
	name := "Novo Nome"
	description := "Nova Descrição"
	address := "Novo Endereço"
	mapUrl := "http://example.com/novo"
	date := time.Now().Add(48 * time.Hour)
	modality := entities.EventModality("virtual")
	cancellationPolicy := "Nova Política"
	participantEditionPolicy := "Nova Edição"
	ticketType := "Regular"
	ticketPrice := 150.0
	ticketQuantity := 20

	validDTO := UpdateEventDTO{
		Name:                     &name,
		Description:              &description,
		Address:                  &address,
		MapUrl:                   &mapUrl,
		Date:                     &date,
		Modality:                 &modality,
		CancellationPolicy:       &cancellationPolicy,
		ParticipantEditionPolicy: &participantEditionPolicy,
		TicketType:               &ticketType,
		TicketPrice:              &ticketPrice,
		TicketQuantity:           &ticketQuantity,
	}

	if err := validDTO.Validate(); err != nil {
		t.Errorf("Validação esperada para UpdateEventDTO válido, mas ocorreu erro: %v", err)
	}

	// Caso inválido: Nome vazio (mínimo 1 caractere esperado)
	invalidName := ""
	validDTO.Name = &invalidName
	if err := validDTO.Validate(); err == nil {
		t.Errorf("Era esperado erro de validação para Nome vazio no UpdateEventDTO")
	}
	validDTO.Name = &name // restaura valor válido

	// Caso inválido: MapUrl inválido
	invalidMapUrl := "invalid_url"
	validDTO.MapUrl = &invalidMapUrl
	if err := validDTO.Validate(); err == nil {
		t.Errorf("Era esperado erro de validação para MapUrl inválido no UpdateEventDTO")
	}
	validDTO.MapUrl = &mapUrl // restaura valor válido

	// Caso inválido: TicketPrice negativo
	negativeTicketPrice := -10.0
	validDTO.TicketPrice = &negativeTicketPrice
	if err := validDTO.Validate(); err == nil {
		t.Errorf("Era esperado erro de validação para TicketPrice negativo no UpdateEventDTO")
	}
	validDTO.TicketPrice = &ticketPrice // restaura valor válido

	// Caso inválido: TicketQuantity zero
	zeroTicketQuantity := 0
	validDTO.TicketQuantity = &zeroTicketQuantity
	if err := validDTO.Validate(); err == nil {
		t.Errorf("Era esperado erro de validação para TicketQuantity <= 0 no UpdateEventDTO")
	}
}
