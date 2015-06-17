package main

import (
	"fmt"
	"github.com/gerschkin/config"
	"github.com/gerschkin/server/rpc"
	"github.com/valyala/gorpc"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/inconshreveable/log15.v2"
	"os"
)

var (
	conf config.Server
	log  log15.Logger
)

func main() {
	// Parse flags
	kingpin.Parse()

	// Create loggers.
	log = log15.New()

	// Try and read the configuration file. If we can't read the file then
	// we exit with exit code 1.
	err := config.ReadServer(*configPath, &conf)
	if err != nil {
		log.Crit("could not read config", "error", err)
		os.Exit(1)
	}

	// Register rpc types that we accept, and figure out the address to use
	// with the rpc server.
	registerTypes()
	addr := fmt.Sprintf("%s:%d", conf.RPC.Host, conf.RPC.Port)

	// Create our TCP RPC server, and serve it. Keep serving it until we
	// receive an error. If we ever get an error then we exit with exit
	// code 2.
	server := gorpc.NewTCPServer(addr, handleRequests)
	if err := server.Serve(); err != nil {
		log.Crit("tpc rpc server stopped", "error", err)
		os.Exit(2)
	}

	defer server.Stop()
}

// We need to register all of the types that our rpc server can
// send and receive. Normal types such as int, string, ... are
// automatically registered and work.
func registerTypes() {
	gorpc.RegisterType(rpc.Request{})
	gorpc.RegisterType(rpc.Response{})
}
