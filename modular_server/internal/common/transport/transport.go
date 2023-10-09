package transport

type HTTPTransport interface {
	Listen(Port string) error
}

type GRPCTransport interface {
	Listen(Port string) error
}
