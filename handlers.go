package main

import (
	"net/http"
	"path/filepath"
)

func (s *server) handleSPA(spaDir string) http.HandlerFunc {
  const spaFile = "index.html"
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(spaDir, spaFile))
	}
}
