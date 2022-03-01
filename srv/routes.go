package app

import (
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() *gin.Engine {
	router := s.router

	router.Static("/css", "static/dist/css")
	router.Static("/js", "static/dist/js")
	router.StaticFile("/favicon.ico", "static/dist/favicon.ico")

	router.GET("/", s.ServeFrontend())
	router.GET("/updates", s.ParseFR())

	return router
}
