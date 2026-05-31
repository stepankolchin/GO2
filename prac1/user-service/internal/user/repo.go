package user

import "errors"

var ErrUserNotFound = errors.New("user not found")

type Repo struct {
	data map[int64]User
}

func NewRepo() *Repo {
	return &Repo{
		data: map[int64]User{
			1: {ID: 1, Name: "Иван Иванов", Email: "ivan@example.com"},
			2: {ID: 2, Name: "Мария Петрова", Email: "maria@example.com"},
			3: {ID: 3, Name: "Алексей Сидоров", Email: "alex@example.com"},
		},
	}
}

func (r *Repo) GetByID(id int64) (User, error) {
	u, ok := r.data[id]
	if !ok {
		return User{}, ErrUserNotFound
	}
	return u, nil
}

func (r *Repo) GetAllUsers() (Users, error) {
	users := make(Users, 0, len(r.data))
	for _, user := range r.data {
		users = append(users, user)
	}
	return users, nil
}

