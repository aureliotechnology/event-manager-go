# Event Manager Go

Projeto de gerenciamento de eventos desenvolvido em Go. Este projeto utiliza MongoDB para persistência e expõe uma API RESTful com endpoints para criação, atualização, consulta e deleção de eventos.

## Características

- **CRUD de Eventos**: Criação, listagem, atualização e deleção de eventos.
- **Validação via DTOs**: Utiliza o [go-playground/validator](https://github.com/go-playground/validator) para validar os dados de entrada.
- **MongoDB**: Persistência utilizando o driver oficial do MongoDB.
- **Gorilla Mux**: Roteamento com [Gorilla Mux](https://github.com/gorilla/mux).
- **Testes Unitários**: Testes escritos para os handlers, DTOs, entidade e repositório.

## Estrutura do Projeto

```
event-manager-go/
├── cmd/
│   └── main.go          # Inicializa o servidor e registra os endpoints.
├── internal/
│   ├── adapters/
│   │   └── httpadapter/ # Handlers HTTP da API.
│   ├── domain/
│   │   ├── dto/         # Objetos de transferência de dados (DTOs) com suas validações.
│   │   └── entities/    # Entidades do domínio (ex.: Event).
│   └── infrastructure/
│       └── persistence/ # Implementação do repositório de eventos (MongoDB).
└── README.md
```

## Instalação

1. **Pré-requisitos**:
   - [Go 1.20+](https://golang.org/dl/)
   - [MongoDB](https://www.mongodb.com/try/download/community) (para testes e execução)
   - (Opcional) [Docker](https://www.docker.com/) para execução do MongoDB em container

2. **Clone o repositório**:

```bash
git clone https://github.com/aureliotechnology/event-manager-go.git
cd event-manager-go
```

3. **Configuração do ambiente**:
Crie um arquivo `.env` na raiz do projeto com as variáveis de ambiente necessárias, por exemplo:

```env
MONGO_URI=mongodb://mongodb:27018
MONGO_DB=eventdb
MONGO_COLLECTION=events
```

## Execução

Para rodar o projeto, execute o comando:

```bash
make build
make docker-up
```

O servidor será iniciado na porta definida (default: 8080).

## Endpoints da API

- **GET /health**
  _Verifica o status do servidor._

- **POST /events**
  _Cria um novo evento._
  Corpo (JSON):
  ```json
  {
      "name": "Nome do evento",
      "description": "Descrição do evento",
      "address": "Endereço do evento",
      "mapUrl": "http://maps.exemplo.com",
      "date": "2025-05-01T10:00:00Z",
      "modality": "presencial", // ou "virtual", "hibrido"
      "cancellationPolicy": "Política de cancelamento",
      "participantEditionPolicy": "Política de edição",
      "ticketType": "VIP",
      "ticketPrice": 150.0,
      "ticketQuantity": 100
  }
  ```

- **GET /events**
  _Lista todos os eventos._

- **GET /events/{id}**
  _Retorna os detalhes de um evento específico._

- **PUT /events/{id}**
  _Atualiza os dados de um evento._
  Corpo (JSON com campos opcionais):
  ```json
  {
      "name": "Nome atualizado",
      "ticketPrice": 200.0
  }
  ```

- **DELETE /events/{id}**
  _Remove um evento._

## Testes

Para executar os testes unitários do projeto, utilize o comando:

```bash
make docker-test
```

Isso executará os testes de DTO, entidade, handlers e repositório.

## Postman

Uma coleção Postman já foi criada para facilitar os testes da API. Importe o arquivo de coleção (formato JSON) no Postman para acessar os endpoints e testar as operações.

## Contribuições

Sinta-se à vontade para enviar _issues_ e _pull requests_.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
