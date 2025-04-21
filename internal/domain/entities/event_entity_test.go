package entities

import (
	"testing"
	"time"
)

func TestNewEvent_WithProvidedID(t *testing.T) {
	id := "123e4567-e89b-12d3-a456-426614174000"
	name := "Evento Teste"
	description := "Descrição de teste"
	address := "Rua Teste, 123"
	mapUrl := "http://maps.example.com"
	date := time.Now().Add(24 * time.Hour)
	modality := Presencial
	cancellationPolicy := "Política de cancelamento"
	participantEditionPolicy := "Política de edição"
	ticketType := "VIP"
	ticketPrice := 150.0
	ticketQuantity := 100

	event := NewEvent(
		name,
		description,
		address,
		mapUrl,
		date,
		modality,
		cancellationPolicy,
		participantEditionPolicy,
		ticketType,
		ticketPrice,
		ticketQuantity,
		id,
	)

	if event.ID != id {
		t.Errorf("Esperado ID %s, obtido %s", id, event.ID)
	}
	if event.Name != name {
		t.Errorf("Esperado Name %s, obtido %s", name, event.Name)
	}
	if event.Description != description {
		t.Errorf("Esperado Description %s, obtido %s", description, event.Description)
	}
	if event.Address != address {
		t.Errorf("Esperado Address %s, obtido %s", address, event.Address)
	}
	if event.MapUrl != mapUrl {
		t.Errorf("Esperado MapUrl %s, obtido %s", mapUrl, event.MapUrl)
	}
	if !event.Date.Equal(date) {
		t.Errorf("Esperado Date %v, obtido %v", date, event.Date)
	}
	if event.Modality != modality {
		t.Errorf("Esperado Modality %s, obtido %s", modality, event.Modality)
	}
	if event.CancellationPolicy != cancellationPolicy {
		t.Errorf("Esperado CancellationPolicy %s, obtido %s", cancellationPolicy, event.CancellationPolicy)
	}
	if event.ParticipantEditionPolicy != participantEditionPolicy {
		t.Errorf("Esperado ParticipantEditionPolicy %s, obtido %s", participantEditionPolicy, event.ParticipantEditionPolicy)
	}
	if event.TicketType != ticketType {
		t.Errorf("Esperado TicketType %s, obtido %s", ticketType, event.TicketType)
	}
	if event.TicketPrice != ticketPrice {
		t.Errorf("Esperado TicketPrice %v, obtido %v", ticketPrice, event.TicketPrice)
	}
	if event.TicketQuantity != ticketQuantity {
		t.Errorf("Esperado TicketQuantity %d, obtido %d", ticketQuantity, event.TicketQuantity)
	}
}

func TestNewEvent_WithoutProvidedID(t *testing.T) {
	// Ao passar uma string vazia para ID, espera-se que um novo UUID seja gerado.
	name := "Evento Sem ID"
	description := "Descrição sem ID"
	address := "Rua Sem ID, 456"
	mapUrl := "http://maps.example.com/semid"
	date := time.Now().Add(48 * time.Hour)
	modality := Virtual
	cancellationPolicy := "Política sem ID"
	participantEditionPolicy := "Política de edição sem ID"
	ticketType := "Standard"
	ticketPrice := 80.0
	ticketQuantity := 50

	event := NewEvent(
		name,
		description,
		address,
		mapUrl,
		date,
		modality,
		cancellationPolicy,
		participantEditionPolicy,
		ticketType,
		ticketPrice,
		ticketQuantity,
		"",
	)

	if event.ID == "" {
		t.Error("Esperado que um novo ID fosse gerado, mas o ID está vazio")
	}
	if event.Name != name {
		t.Errorf("Esperado Name %s, obtido %s", name, event.Name)
	}
	if event.Description != description {
		t.Errorf("Esperado Description %s, obtido %s", description, event.Description)
	}
	if event.Address != address {
		t.Errorf("Esperado Address %s, obtido %s", address, event.Address)
	}
	if event.MapUrl != mapUrl {
		t.Errorf("Esperado MapUrl %s, obtido %s", mapUrl, event.MapUrl)
	}
	if !event.Date.Equal(date) {
		t.Errorf("Esperado Date %v, obtido %v", date, event.Date)
	}
	if event.Modality != modality {
		t.Errorf("Esperado Modality %s, obtido %s", modality, event.Modality)
	}
	if event.CancellationPolicy != cancellationPolicy {
		t.Errorf("Esperado CancellationPolicy %s, obtido %s", cancellationPolicy, event.CancellationPolicy)
	}
	if event.ParticipantEditionPolicy != participantEditionPolicy {
		t.Errorf("Esperado ParticipantEditionPolicy %s, obtido %s", participantEditionPolicy, event.ParticipantEditionPolicy)
	}
	if event.TicketType != ticketType {
		t.Errorf("Esperado TicketType %s, obtido %s", ticketType, event.TicketType)
	}
	if event.TicketPrice != ticketPrice {
		t.Errorf("Esperado TicketPrice %v, obtido %v", ticketPrice, event.TicketPrice)
	}
	if event.TicketQuantity != ticketQuantity {
		t.Errorf("Esperado TicketQuantity %d, obtido %d", ticketQuantity, event.TicketQuantity)
	}
}
