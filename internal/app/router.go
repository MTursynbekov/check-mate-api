package app

func (s *Server)Route(){
	s.app.Post("/message", s.CreateMessage)
	s.app.Get("/messages/:chatId", s.GetMessages)
	s.app.Post("/chat", s.CreateChat)
}