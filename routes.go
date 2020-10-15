package main

import (
	"net/http"
	"os"
)

func (s *server) routes() {
	r := s.router

	// Check for SPA
	spaDir := "./client/public/"
	_, err := os.Stat(spaDir)
	if err != nil {
		s.panic("'/client/public/' not detected!", err)
	}

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(spaDir))))
	r.HandleFunc("/", s.handleSPA(spaDir))

}
