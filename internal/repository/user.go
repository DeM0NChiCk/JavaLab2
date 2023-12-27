package repository

import "HomeworkWebProject/internal/model"

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) GetAll() (users []*model.User) {
	users = []*model.User{
		{
			ID:        1,
			FirstName: "Dima",
			LastName:  "Toropov",
		},
		{
			ID:        2,
			FirstName: "Nikita",
			LastName:  "Chivilev",
		},
	}

	return
}
