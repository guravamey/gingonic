package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/guravamey/gingonic/entity"
	"github.com/guravamey/gingonic/service"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) entity.User
	Update(id string, ctx *gin.Context) entity.User
	Delete(id string)
}

type controller struct {
	services service.UserService
}

func New() UserController {
	return &controller{services: service.New()}
}

func (c controller) FindAll() []entity.User {
	return c.services.FindAll()
}

func (c controller) Save(ctx *gin.Context) entity.User {
	var userDetail entity.User
	userDetail.ID = uuid.New().String()
	if err := ctx.ShouldBindJSON(&userDetail); err != nil {
		ctx.JSON(422, gin.H{"error": true, "message": "Invalid Request Body"})
		return entity.User{}
	}
	return c.services.Save(userDetail)
}

func (c controller) Update(id string, ctx *gin.Context) entity.User {
	var userDetail entity.User
	if err := ctx.ShouldBindJSON(&userDetail); err != nil {
		ctx.JSON(422, gin.H{"error": true, "message": "Invalid Request Body"})
		return entity.User{}
	}
	return c.services.Update(id, userDetail)
}

func (c controller) Delete(id string) {
	c.services.Delete(id)
}
