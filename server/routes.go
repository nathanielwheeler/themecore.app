package server

func (s *server) routes() {
	r := s.router

	r.HandleFunc("/", s.handleSPA())

}
