package main

import (
  "encoding/json"
	"log"
	"net/http"
	"fmt"
  "os"
  
  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Always run first
	s := newServer()

	s.logger.Printf("Now listening on %s...\n", s.vars.Port)
	http.ListenAndServe(s.vars.Port, s.router)
	return nil
}

type server struct {
	router *mux.Router
	logger *log.Logger
  vars   varConfig
  db *gorm.DB
}

func newServer() *server {
  s := &server{
    router: mux.NewRouter(),
		logger: log.New(os.Stdout, "themecore.app: ", log.Lshortfile),
	}
  cfg := s.loadConfig()
  s.vars = cfg.Vars

  s.setupDB(cfg.DB)
  defer s.db.Close()
  s.autoMigrate()

	s.routes()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}



// Helpers

func (s *server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			s.logErr("JSON Encoding failed", err)
		}
	}
}

func (s *server) logErr(message string, err interface{}) {
	s.logger.Fatalf("%s\n\t%s", message, err)
}

func (s *server) logMsg(message string) {
	s.logger.Println(message)
}

func (s *server) panic(message string, err interface{}) {
	s.logger.Panicf("%s\n\t%s", message, err)
}
