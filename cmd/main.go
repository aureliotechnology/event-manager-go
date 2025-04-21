package main

import (
	"log"
	"net/http"
	"os"

	"event-manager-go/internal/adapters/httpadapter"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega as variáveis do .env
	if err := godotenv.Load(); err != nil {
		log.Println("Nenhum .env encontrado, utilizando variáveis de ambiente do sistema.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Cria um novo router Gorilla Mux
	router := mux.NewRouter()

	// Endpoint para Health Check
	router.HandleFunc("/health", httpadapter.HealthHandler).Methods("GET")

	// Endpoint para criar um novo evento (POST /events)
	router.HandleFunc("/events", httpadapter.CreateEventHandler).Methods("POST")

	// Endpoint para buscar todos os eventos (GET /events)
	router.HandleFunc("/events", httpadapter.FindAllEventHandler).Methods("GET")

	// Endpoint para atualizar um evento (PUT /events/{id})
	router.HandleFunc("/events/{id}", httpadapter.UpdateEventHandler).Methods("PUT")
	// Endpoint para deletar um evento (DELETE /events/{id})
	router.HandleFunc("/events/{id}", httpadapter.DeleteEventHandler).Methods("DELETE")
	// Endpoint para buscar um evento específico (GET /events/{id}
	router.HandleFunc("/events/{id}", httpadapter.GetEventHandler).Methods("GET")

	log.Printf("Servidor rodando na porta :%s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
