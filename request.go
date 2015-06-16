package main

import (
	"github.com/gerschkin/server/rpc"
)

//
func handleRequests(clientAddr string, request interface{}) interface{} {

	return rpc.Response{
		Request: request.(rpc.Request),
	}
}
