package config

import "os"

type Server struct {
	Port string
}

func (s *Server) Load() *Server {
	s.Port = os.Getenv("SERVER_PORT")

	if s.Port == "" {
		s.Port = "8080"
	}
	return s
}
