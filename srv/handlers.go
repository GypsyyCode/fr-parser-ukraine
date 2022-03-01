package app

import (
	"fr.de-parser/parser"
	"github.com/gin-gonic/gin"
)

// func (s *Server) ServeStatic() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.File("../static/frontend/dist")
// 	}
// }

// func (s *Server) ServeAdmin() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.File("../static/admin/dist/index.html")
// 	}
// }

func (s *Server) ParseFR() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content.Type", "application/json")
		url := "https://www.fr.de/politik/ukraine-konflikt-krieg-russland-kiew-kampf-invasion-hauptstadt-news-zr-91366920.html"
		res := parser.ParseFR(url)

		if res == nil {
			c.JSON(500, gin.H{"message": "error"})
			return
		}

		c.JSON(200, gin.H{"message": res})
	}
}

func (s *Server) ServeFrontend() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.File("static/dist/index.html")
	}
}
