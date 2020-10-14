package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type server struct {
	router  *mux.Router
	logger  *log.Logger
	spaDir string
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
		logger: log.New(os.Stdout, "themecore.app: ", log.Lshortfile),
		spaDir: "./client/public/",
	}
	// Check for client
	_, err := os.Stat(s.spaDir)
	if err != nil {
		s.logger.Panicf("'/client/public/' not detected!\n\t%s\n", err)
	}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Run starts the server and returns an error to main if anything goes wrong.
func Run() error {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	s := newServer()

	s.logger.Printf("Now listening on %s...\n", port)
	http.ListenAndServe(port, s.router)
	return nil
}
