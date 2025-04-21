package persistence

import (
	"context"
	"time"

	"event-manager-go/internal/domain/dto"
	"event-manager-go/internal/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoEventRepository implementa as operações de persistência para a entidade Event.
type MongoEventRepository struct {
	collection *mongo.Collection
}

// NewMongoEventRepository cria e retorna uma nova instância de MongoEventRepository.
func NewMongoEventRepository(uri, dbName, collName string) (*MongoEventRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collName)
	return &MongoEventRepository{collection: collection}, nil
}

// CreateEventHandler insere um novo evento no MongoDB.
func (r *MongoEventRepository) Create(ctx context.Context, event *entities.Event) error {
	_, err := r.collection.InsertOne(ctx, event)
	return err
}

// FindAllEventHandler retorna um cursor com todos os eventos.
func (r *MongoEventRepository) FindAllEventHandler(ctx context.Context, filter bson.M) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return r.collection.Find(ctx, filter)
}

// FindByID busca e retorna um único evento pelo seu ID.
func (r *MongoEventRepository) FindByID(ctx context.Context, id string) (*entities.Event, error) {
	var event entities.Event
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// Update atualiza um evento existente com os campos não-nulos fornecidos em updateDTO.
func (r *MongoEventRepository) Update(ctx context.Context, id string, updateDTO dto.UpdateEventDTO) (*entities.Event, error) {
	updateFields := bson.M{}

	if updateDTO.Name != nil {
		updateFields["name"] = *updateDTO.Name
	}
	if updateDTO.Description != nil {
		updateFields["description"] = *updateDTO.Description
	}
	if updateDTO.Address != nil {
		updateFields["address"] = *updateDTO.Address
	}
	if updateDTO.MapUrl != nil {
		updateFields["mapUrl"] = *updateDTO.MapUrl
	}
	if updateDTO.Date != nil {
		updateFields["date"] = *updateDTO.Date
	}
	if updateDTO.Modality != nil {
		updateFields["modality"] = *updateDTO.Modality
	}
	if updateDTO.CancellationPolicy != nil {
		updateFields["cancellationPolicy"] = *updateDTO.CancellationPolicy
	}
	if updateDTO.ParticipantEditionPolicy != nil {
		updateFields["participantEditionPolicy"] = *updateDTO.ParticipantEditionPolicy
	}
	if updateDTO.TicketType != nil {
		updateFields["ticketType"] = *updateDTO.TicketType
	}
	if updateDTO.TicketPrice != nil {
		updateFields["ticketPrice"] = *updateDTO.TicketPrice
	}
	if updateDTO.TicketQuantity != nil {
		updateFields["ticketQuantity"] = *updateDTO.TicketQuantity
	}

	if len(updateFields) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedEvent entities.Event
	err := r.collection.FindOneAndUpdate(ctx, bson.M{"id": id}, bson.M{"$set": updateFields}, opts).Decode(&updatedEvent)
	if err != nil {
		return nil, err
	}
	return &updatedEvent, nil
}

// Delete remove um evento pelo seu ID.
func (r *MongoEventRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"id": id})
	return err
}
