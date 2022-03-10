package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fr.de-parser/parser"
	"github.com/gin-gonic/gin"
)

type GoogleNewsFeed struct {
	Feed     interface{} `json:"feed"`
	Articles interface{} `josn:"articles"`
}

func (s *Server) ParseFR() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content.Type", "application/json")
		url := "https://www.fr.de/politik/ukraine-krieg-russland-kiew-invasion-kaempfe-militaer-wladimir-putin-wolodymyr-selenskyj-konflikt-news-ticker-zr-91391160.html"
		res := parser.ParseFR(url)

		if res == nil {
			c.JSON(500, gin.H{"message": "error"})
			return
		}

		c.JSON(200, gin.H{"message": res})
	}
}

func (s *Server) GetGoogleNewsFeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := "https://google-news.p.rapidapi.com/v1/search?q=Ukraine&country=DE&lang=de"

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("x-rapidapi-host", "google-news.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", "f3e34c28a0msh4bdecec013b6116p12865djsn666384434311")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		var newsFeed GoogleNewsFeed

		if err := json.Unmarshal(body, &newsFeed); err != nil {
			fmt.Printf("%+v\n", err)
			c.JSON(500, gin.H{"message": "unabel to request data"})
			return
		}

		c.JSON(200, gin.H{"message": newsFeed})

		fmt.Println(res)
		fmt.Println(string(body))

	}
}

func (s *Server) ServeFrontend() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.File("static/dist/index.html")
	}
}
