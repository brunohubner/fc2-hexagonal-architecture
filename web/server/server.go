package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/brunohubner/fc2-hexagonal-architecture/src/application"
	"github.com/brunohubner/fc2-hexagonal-architecture/web/handlers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.IProductService
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handlers.MakeProductHandler(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           nil,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
