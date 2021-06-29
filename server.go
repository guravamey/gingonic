package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guravamey/gingonic/controller"
)

//Local host 5000, get,put, post and delete. ID as path param for put and delete.
//POST and PUT requires body. Body contains Name and Age

func main() {

	usercontroller := controller.New()
	server := gin.Default()
	userRoutes := server.Group("/users")
	{
		userRoutes.GET("/", func(ctx *gin.Context) { ctx.JSON(200, usercontroller.FindAll()) })
		userRoutes.POST("/", func(ctx *gin.Context) { ctx.JSON(200, usercontroller.Save(ctx)) })
		userRoutes.PUT("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			ctx.JSON(200, usercontroller.Update(id, ctx))
		})
		userRoutes.DELETE("/:id", func(ctx *gin.Context) {
			usercontroller.Delete(ctx.Param("id"))
			ctx.JSON(200, gin.H{"message": "Success"})
		})
	}
	server.Run(":5000")

}
