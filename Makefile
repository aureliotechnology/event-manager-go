.PHONY: build run docker-up docker-down clean test

# Compila a aplicação
build:
	@echo "Compilando a aplicação..."
	go build -o event-manager-go ./cmd/main.go

# Executa a aplicação localmente (após compilar)
run: build
	@echo "Executando a aplicação..."
	./event-manager-go

# Inicia a aplicação via Docker Compose
docker-up:
	@echo "Subindo os containers com Docker Compose..."
	docker-compose up --build

# Para os containers do Docker Compose
docker-down:
	@echo "Parando os containers..."
	docker-compose down

# Remove os binários gerados
clean:
	@echo "Limpando o projeto..."
	rm -f event-manager-go

# Executa os testes
test:
	@echo "Executando testes..."
	go test ./...
