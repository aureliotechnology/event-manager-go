package httpadapter

import (
	"fmt"
	"net/http"

	"event-manager-go/internal/application"
)

// HelloHandler é o adaptador HTTP que expõe o endpoint.
type HelloHandler struct {
	helloService application.HelloService
}

// NewHelloHandler cria um novo handler a partir de um serviço.
func NewHelloHandler(service application.HelloService) *HelloHandler {
	return &HelloHandler{
		helloService: service,
	}
}

// ServeHTTP lida com as requisições HTTP.
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := h.helloService.SayHello()
	fmt.Fprintln(w, message)
}
