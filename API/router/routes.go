package router

import (
	"BackendGo/middlewares"
	"BackendGo/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApplyRoutes(s *server.Server) error {
	s.Router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	s.Router.POST("/login", loginUser(s))
	s.Router.GET("/query", middlewares.JwtAuth(), queries(s))
	s.Router.GET("/query/:exporter", middlewares.JwtAuth(), queryExporter(s))
	s.Router.GET("/JWTAuthTest", middlewares.JwtAuth(), func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	return nil
}
