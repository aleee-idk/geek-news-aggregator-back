package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type _new struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var news = []_new{
	{Id: "1", Title: "Lisa new Album", Text: "The artist LiSa Creates a new album!!"},
	{Id: "2", Title: "New TES Skyrim port", Text: "A new port of the bethesda bestselers has come out, now in your friedge"},
	{Id: "3", Title: "Boku no Hero S6", Text: "The new season for Boku no Hero is now in development"},
}

func main() {
	router := gin.Default()
	router.GET("/news", getNews)

	router.Run("localhost:3000")
}

func getNews(c *gin.Context) {
	c.IndentedJSON(200, news)
	fmt.Println(news)
}
