package persistence

import (
    "context"
    "testing"
    "time"

    "event-manager-go/internal/domain/dto"
    "event-manager-go/internal/domain/entities"

    "go.mongodb.org/mongo-driver/bson"
)

// setupRepository cria um repositório de teste e limpa a coleção para assegurar um ambiente limpo.
func setupRepository(t *testing.T) *MongoEventRepository {
    repo, err := NewMongoEventRepository("mongodb://localhost:27017", "testdb", "testevents")
    if err != nil {
        t.Fatalf("Erro ao conectar com o MongoDB: %v", err)
    }

    // Limpa a coleção antes de cada teste.
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := repo.collection.Drop(ctx); err != nil {
        t.Fatalf("Erro ao limpar a coleção: %v", err)
    }

    return repo
}

// createTestEvent cria uma instância de evento para os testes.
func createTestEvent() *entities.Event {
    return &entities.Event{
        ID:                       "test-" + time.Now().Format("20060102150405"),
        Name:                     "Evento de Teste",
        Description:              "Descrição de teste",
        Address:                  "Rua Teste, 123",
        MapUrl:                   "http://example.com/test",
        Date:                     time.Now().Add(24 * time.Hour),
        Modality:                 entities.Presencial,
        CancellationPolicy:       "Política de Cancelamento",
        ParticipantEditionPolicy: "Política de Edição",
        TicketType:               "VIP",
        TicketPrice:              100.0,
        TicketQuantity:           50,
    }
}

func TestCreateAndFindByID(t *testing.T) {
    repo := setupRepository(t)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Cria um evento de teste.
    event := createTestEvent()
    if err := repo.Create(ctx, event); err != nil {
        t.Fatalf("Erro ao criar evento: %v", err)
    }

    // Busca o evento pelo ID.
    found, err := repo.FindByID(ctx, event.ID)
    if err != nil {
        t.Fatalf("Erro ao buscar evento por ID: %v", err)
    }
    if found.ID != event.ID {
        t.Errorf("Evento encontrado com ID %s; esperado %s", found.ID, event.ID)
    }
}

func TestFindAllEventHandler(t *testing.T) {
    repo := setupRepository(t)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Cria e insere dois eventos.
    event1 := createTestEvent()
    event2 := createTestEvent()
    if err := repo.Create(ctx, event1); err != nil {
        t.Fatalf("Erro ao criar evento1: %v", err)
    }
    if err := repo.Create(ctx, event2); err != nil {
        t.Fatalf("Erro ao criar evento2: %v", err)
    }

    // Busca todos os eventos.
    cursor, err := repo.FindAllEventHandler(ctx, bson.M{})
    if err != nil {
        t.Fatalf("Erro ao buscar todos os eventos: %v", err)
    }
    defer cursor.Close(ctx)

    var events []entities.Event
    if err := cursor.All(ctx, &events); err != nil {
        t.Fatalf("Erro ao decodificar eventos: %v", err)
    }
    if len(events) != 2 {
        t.Errorf("Número de eventos retornados = %d; esperado 2", len(events))
    }
}

func TestUpdate(t *testing.T) {
    repo := setupRepository(t)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Cria um evento.
    event := createTestEvent()
    if err := repo.Create(ctx, event); err != nil {
        t.Fatalf("Erro ao criar evento: %v", err)
    }

    // Define um updateDTO para mudar alguns campos.
    newName := "Evento Atualizado"
    newTicketPrice := 150.0
    updateDTO := dto.UpdateEventDTO{
        Name:        &newName,
        TicketPrice: &newTicketPrice,
    }

    updated, err := repo.Update(ctx, event.ID, updateDTO)
    if err != nil {
        t.Fatalf("Erro ao atualizar evento: %v", err)
    }
    if updated.Name != newName {
        t.Errorf("Nome atualizado incorreto: esperado %s, obtido %s", newName, updated.Name)
    }
    if updated.TicketPrice != newTicketPrice {
        t.Errorf("TicketPrice atualizado incorreto: esperado %v, obtido %v", newTicketPrice, updated.TicketPrice)
    }
}

func TestDelete(t *testing.T) {
    repo := setupRepository(t)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Cria um evento para ser deletado.
    event := createTestEvent()
    if err := repo.Create(ctx, event); err != nil {
        t.Fatalf("Erro ao criar evento: %v", err)
    }

    // Deleta o evento.
    if err := repo.Delete(ctx, event.ID); err != nil {
        t.Fatalf("Erro ao deletar evento: %v", err)
    }

    // Tenta buscar o evento deletado – espera que não seja encontrado.
    _, err := repo.FindByID(ctx, event.ID)
    if err == nil {
        t.Errorf("Esperava erro ao buscar evento deletado, mas o evento foi encontrado")
    }
}
