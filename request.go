package main

import (
	"github.com/gerschkin/server/rpc"
)

// The unexported func authenticated takes the request and
// tries to authenticate it. If the request does get
// authenticated then we can continue, otherwise we
// must return a not authorized response.
func authenticated(request rpc.Request) bool {
	// TODO
	// - authenticate
	return true
}

// The unexported func handleRequests handles all incoming tcp
// rpc requests. This func also tries to authenticate clients,
// and responds accordingly to the requests.
func handleRequests(clientAddr string, request interface{}) interface{} {
	// Cast request from an interface to rpc.Request
	req := request.(rpc.Request)

	// Setup request logger for this client
	reqlog := log.New("clientAddr", clientAddr)
	reqlog.Info("received request", "request", req)

	// Now we want to try and authenticate the request. If we can
	// not authenticate the request then we want to return the
	// response as RESPONSE_NOT_AUTHORIZED.
	if !authenticated(req) {
		reqlog.Info("unauthorized request", "request", req)
		return rpc.ResponseNotAuthorized(req)
	} else {
		reqlog.Info("request authorized", "request", req)
	}

	// Now we're at the dispatcher. From here we figure out what
	// type of request we've been given and act accordingly.
	switch req.Type {
	case rpc.REQUEST_PING:
		reqlog.Info("ping received", "request", req)
		return pong(req)
	}

	// If we get this far then we have an invalid request.
	reqlog.Info("invalid request", "request", req)
	return invalidRequest(req)
}

// The unexported func pong just sends a response back to the
// client with the original request and a RESPONSE_OK. This
// is to be used if the client requests a pong (by sending
// a ping).
func pong(request rpc.Request) rpc.Response {
	return rpc.Response{
		Request: request,
		Status:  rpc.RESPONSE_OK,
	}
}

// The unexported func invalidRequest just sends a resposne back
// to the client with the original request and the status
// RESPONSE_REQUEST_TYPE_INVALID. This is to be used if
// a client sends an unknown request.
func invalidRequest(request rpc.Request) rpc.Response {
	return rpc.Response{
		Request: request,
		Status:  rpc.RESPONSE_REQUEST_TYPE_INVALID,
	}
}
