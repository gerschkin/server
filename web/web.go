package web

import (
	"fmt"
	"github.com/gerschkin/config"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var router *httprouter.Router

// The exported func Serve sets up our http server and adds
// routes/handlers.
func Serve(conf config.Server) {
	// Create our router
	router = httprouter.New()

	// Add routes
	route()

	// Figure out the address to host the web on
	addr := fmt.Sprintf("%s:%d", conf.Web.Host, conf.Web.Port)
	http.ListenAndServe(addr, router)
}
