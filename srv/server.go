package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router            *gin.Engine
	/* Services */
}

func NewServer(router *gin.Engine,
	/* services */) *Server {
	return &Server{
		router:            router,
		/* Services */
	}
}

func (s *Server) Run() error {
	r := s.Routes()

	err := r.Run(":80")

	if err != nil {
		log.Printf("Server - there was an error calling Run on rounter: %v", err)
		return err
	}
	return nil
}
