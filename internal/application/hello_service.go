package application

// HelloService define o contrato de serviço da aplicação.
type HelloService interface {
    SayHello() string
}

// helloServiceImpl é uma implementação de HelloService.
type helloServiceImpl struct{}

// SayHello retorna a mensagem "Hello World".
func (s *helloServiceImpl) SayHello() string {
    return "Hello World"
}

// NewHelloService cria uma instância de HelloService.
func NewHelloService() HelloService {
    return &helloServiceImpl{}
}
