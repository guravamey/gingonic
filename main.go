package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

//Local host 5000, get,put, post and delete. ID as path param for put and delete. 
//POST and PUT requires body. Body contains Name and Age

func main() {
	r := gin.Default()
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", getUsers)
		userRoutes.POST("/", createUser)
		userRoutes.PUT("/:id", editUser)
		userRoutes.DELETE("/:id", deleteUser)
	}

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users []User      //nil
var Users2 = []User{} //empty

func getUsers(c *gin.Context) {
	c.JSON(200, Users)
}

func createUser(c *gin.Context) {
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{"error": true, "message": "Invalid Request Body"})
		return
	}
	reqBody.ID = uuid.New().String()
	Users = append(Users, reqBody)
	c.JSON(200, gin.H{"error": false})

}

func editUser(c *gin.Context) {
	id := c.Param("id")
	println(id)
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{"error": true, "message": "Invalid Request Body"})
		return
	}
	for i, u:= range Users{
		if u.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age
			
			c.JSON(200, gin.H{"error": false})
			return
		}
	}
	
	c.JSON(404, gin.H{"error": true, "message": "ID not found"})
	
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	println(id)
	for i, u:= range Users{
		if u.ID == id {
			Users = append(Users[:i], Users[i+1:]...)
			c.JSON(200, gin.H{"error": false})
			return
		}
	}
	
	c.JSON(404, gin.H{"error": true, "message": "ID not found"})
	
}
