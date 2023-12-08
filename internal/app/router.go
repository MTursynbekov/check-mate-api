package app

func (s *Server)Route(){
	api := s.app.Group("/api")

	api.Post("/message", s.CreateMessage)
	api.Get("/messages/:chatId", s.GetMessages)

	api.Post("/chat", s.CreateChat)

	api.Post("/contact", s.CreateContact)

}