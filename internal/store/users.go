package store

import "check-mate/internal/model"

func (s *store) CreateUser(user *model.User) (uint, error) {
	query := `
	insert into users (username, password, phone)
	values ($1, $2, $3)
	returning id`

	var id uint
	err := s.db.QueryRow(query, user.Username, user.Password, user.Phone).Scan(&id)
	return id, err
}

func (s *store) GetUser(username string) (*model.User, error) {
	user := new(model.User)
	query := `
	select id, username, password from users
	where username = $1`
	err := s.db.Get(&user, query)

	return user, err
}
