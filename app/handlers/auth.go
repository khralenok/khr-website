package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// This function render an HTML for auth page
func ShowAuth(authType string, c *gin.Context) {
	switch authType {
	case "login":
		c.HTML(http.StatusOK, "auth.html", gin.H{
			"title":   "Khralenok - Login",
			"heading": "Login to your account",
			authType:  true,
		})
	case "signin":
		c.HTML(http.StatusOK, "auth.html", gin.H{
			"title":   "Khralenok - Signin",
			"heading": "Create new account",
			authType:  true,
		})
	case "logout":
		c.HTML(http.StatusOK, "auth.html", gin.H{
			"title":  "Khralenok - Logout",
			authType: true,
		})
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong authType value",
		})
		return
	}
}
