package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guravamey/gingonic/controller"
	"github.com/guravamey/gingonic/middleware"
	"io"
	"net/http"
	"os"
)

//Local host 5000, get,put, post and delete. ID as path param for put and delete.
//POST and PUT requires body. Body contains Name and Age

func main() {

	usercontroller := controller.New()
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	userRoutes := server.Group("/users")
	{
		userRoutes.GET("/", func(ctx *gin.Context) { ctx.JSON(200, usercontroller.FindAll()) })
		userRoutes.POST("/", func(ctx *gin.Context) {
			save, err := usercontroller.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "true", "message": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, save)
			}
		})
		userRoutes.PUT("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			save, err := usercontroller.Update(id, ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "true", "message": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, save)
			}
		})
		userRoutes.DELETE("/:id", func(ctx *gin.Context) {
			usercontroller.Delete(ctx.Param("id"))
			ctx.JSON(200, gin.H{"message": "Success"})
		})
	}
	server.Run(":5000")

}

//Writes logs in gin.log file
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
