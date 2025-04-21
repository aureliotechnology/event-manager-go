package httpadapter

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"

	"event-manager-go/internal/domain/dto"
	"event-manager-go/internal/domain/entities"
	"event-manager-go/internal/infrastructure/persistence"

	"github.com/google/uuid"
)

// CreateEventHandler lida com a criação de um novo evento via POST /events.
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	// Decodifica o corpo da requisição para o DTO de criação.
	var createDTO dto.CreateEventDTO
	if err := json.NewDecoder(r.Body).Decode(&createDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Valida os dados do DTO.
	if err := createDTO.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Converte o DTO para a entidade Event.
	event := entities.Event{
		ID:                       uuid.NewString(), // ou gere um ID novo, se necessário
		Name:                     createDTO.Name,
		Description:              createDTO.Description,
		Address:                  createDTO.Address,
		MapUrl:                   createDTO.MapUrl,
		Date:                     createDTO.Date,
		Modality:                 createDTO.Modality,
		CancellationPolicy:       createDTO.CancellationPolicy,
		ParticipantEditionPolicy: createDTO.ParticipantEditionPolicy,
		TicketType:               createDTO.TicketType,
		TicketPrice:              createDTO.TicketPrice,
		TicketQuantity:           createDTO.TicketQuantity,
	}

	// Instancia o repositório.
	repo, err := persistence.NewMongoEventRepository(
		os.Getenv("MONGO_URI"),
		os.Getenv("MONGO_DB"),
		os.Getenv("MONGO_COLLECTION"),
	)
	if err != nil {
		http.Error(w, "Erro ao conectar com banco de dados", http.StatusInternalServerError)
		return
	}

	// Insere a entidade gerada no banco de dados.
	if err := repo.Create(r.Context(), &event); err != nil {
		log.Println(err)
		http.Error(w, "Erro ao criar evento", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "ID do evento é obrigatório", http.StatusBadRequest)
		return
	}

	// Decodifica o corpo da requisição no DTO de update
	var updateDTO dto.UpdateEventDTO
	if err := json.NewDecoder(r.Body).Decode(&updateDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Valida os dados do update DTO
	if err := updateDTO.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo, err := persistence.NewMongoEventRepository(
		os.Getenv("MONGO_URI"),
		os.Getenv("MONGO_DB"),
		os.Getenv("MONGO_COLLECTION"),
	)
	if err != nil {
		http.Error(w, "Erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}

	updatedEvent, err := repo.Update(r.Context(), id, updateDTO)
	if err != nil {
		http.Error(w, "Erro ao atualizar evento", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedEvent); err != nil {
		log.Println("Erro ao codificar a resposta:", err)
	}
}

func FindAllEventHandler(w http.ResponseWriter, r *http.Request) {
	repo, err := persistence.NewMongoEventRepository(
		os.Getenv("MONGO_URI"),
		os.Getenv("MONGO_DB"),
		os.Getenv("MONGO_COLLECTION"),
	)
	if err != nil {
		http.Error(w, "Erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}

	cursor, err := repo.FindAllEventHandler(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, "Erro ao buscar eventos", http.StatusInternalServerError)
		return
	}

	var events []entities.Event
	if err := cursor.All(r.Context(), &events); err != nil {
		http.Error(w, "Erro ao decodificar eventos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

// GetEventHandler lida com a busca de um único evento via GET /events/{id}.
func GetEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "ID do evento é obrigatório", http.StatusBadRequest)
		return
	}

	repo, err := persistence.NewMongoEventRepository(
		os.Getenv("MONGO_URI"),
		os.Getenv("MONGO_DB"),
		os.Getenv("MONGO_COLLECTION"),
	)
	if err != nil {
		http.Error(w, "Erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}

	event, err := repo.FindByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Erro ao buscar evento ou evento não encontrado", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

// DeleteEventHandler lida com a remoção de um evento via DELETE /events/{id}.
func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "ID do evento é obrigatório", http.StatusBadRequest)
		return
	}

	repo, err := persistence.NewMongoEventRepository(
		os.Getenv("MONGO_URI"),
		os.Getenv("MONGO_DB"),
		os.Getenv("MONGO_COLLECTION"),
	)
	if err != nil {
		http.Error(w, "Erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}

	if err := repo.Delete(r.Context(), id); err != nil {
		http.Error(w, "Erro ao deletar evento", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
