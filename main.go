package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/quotes", getMusicQuotes)

	engine.Run()
}

func getMusicQuotes(c *gin.Context) {
	url := "https://api.quotable.io/quotes/random"
	httpClient := http.Client{}
	request, _ := http.NewRequest("GET", url, nil)

	response, _ := httpClient.Do(request)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	s := string(body)
	fmt.Println(s)

	var quotes []Quote
	json.Unmarshal(body, &quotes)
	fmt.Printf("result : %v", quotes)
	c.JSON(http.StatusOK, quotes)
}

type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}
