package controllers

import (
	"github.com/gin-gonic/gin"
	"goserv/models"
	"net/http"
)

func MeEndpoint(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	typedUser := user.(models.User)
	c.JSON(http.StatusOK, gin.H{
		"username":  typedUser.Username,
		"fullName":  typedUser.FullName,
		"email":     typedUser.Email,
		"firstName": typedUser.FirstName,
		"lastName":  typedUser.LastName,
	})
}
