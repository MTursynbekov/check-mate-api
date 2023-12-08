package app

func (s *Server) router() {
	s.app.Post("/signup", s.SignupHandler)
	s.app.Post("/signin", s.SigninHandler)
}
