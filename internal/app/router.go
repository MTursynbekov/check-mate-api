package app

func (s *Server) Route() {
	s.app.Post("/signup", s.SignupHandler)
	s.app.Post("/signin", s.SigninHandler)

	api := s.app.Group("/api")

	api.Post("/message", s.CreateMessage)
	api.Get("/messages/:chatId", s.GetMessages)

	api.Post("/chat", s.CreateChat)
	api.Get("/users/:userId/contacts", s.GetContactsHandler)
	api.Get("/users/:userId/contacts/:contactId", s.GetContactHandler)

	api.Post("/contacts", s.CreateContact)

}
