package main

import (
	"fmt"
	"os"

	app "fr.de-parser/srv"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("this is the startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	server := app.NewServer(router /* ,services */)

	err := server.Run()
	// router := srv.setupRouter()
	if err != nil {
		return err
	}
	return nil

}
