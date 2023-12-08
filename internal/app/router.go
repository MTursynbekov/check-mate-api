package app

func (s *Server)Route(){
	s.app.Post("/message", s.CreateMessage)
}