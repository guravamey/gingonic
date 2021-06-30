package middleware

import "github.com/gin-gonic/gin"

//Adds a basic authentication middleware. Map represents the allowed ID's and Pass combination
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{"gin": "gonic"})
}
