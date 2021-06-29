package service

import (
	"github.com/guravamey/gingonic/entity"
)

type UserService interface {
	Save(user entity.User) entity.User
	FindAll() []entity.User
	Delete(id string)
	Update(id string, user entity.User) entity.User
}

type userService struct {
	users []entity.User
}

func New() UserService {
	return &userService{}
}

func (service *userService) Save(user entity.User) entity.User {
	service.users = append(service.users, user)
	return user

}

func (service *userService) FindAll() []entity.User {
	return service.users
}

func (service *userService) Delete(id string) {
	for i, u := range service.users {
		if u.ID == id {
			service.users = append(service.users[:i], service.users[i+1:]...)
		}
	}
}

func (service *userService) Update(id string, user entity.User) entity.User {
	for i, u := range service.users {
		if u.ID == id {
			service.users[i].Name = user.Name
			service.users[i].Age = user.Age
			return service.users[i]
		}
	}
	return entity.User{}
}
