package router

import (
	"BackendGo/auth"
	"BackendGo/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//TODO: fix Error codes
func loginUser(s *server.Server) func(*gin.Context) {
	return func(c *gin.Context) {
		login := c.Query("login")
		pass := c.Query("pass")

		if login == "" || pass == "" {
			c.JSON(http.StatusBadRequest, fmt.Errorf("invalid credentials"))
			return
		}
		user, err := s.Controller.GetUser(login, pass)
		if err != nil {
			c.JSON(http.StatusBadRequest, user)
			return
		}
		token, err := auth.CreateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, token)
			return
		}
		c.String(http.StatusOK, token)
	}
}
