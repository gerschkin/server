package rpc

const (
	REQUEST_PING = 0
)

// The type Request is a struct containing an rpc request.
// Requests are sent to the server by the client.
type Request struct {
	// Type is an integer that contains the type of request
	// that is being sent from the client to the server.
	Type int
}
