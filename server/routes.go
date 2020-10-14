package server

import (
	"net/http"
	"path/filepath"
)

func (s *server) routes() {
	r := s.router

	r.HandleFunc("/", s.handleSPA())

}

func (s *server) handleSPA() http.HandlerFunc {
  const spaFile = "index.html"
	return func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(s.spaDir))

		http.ServeFile(w, r, filepath.Join(s.spaDir, spaFile))
	}
}
