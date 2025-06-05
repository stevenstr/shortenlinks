package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	gin "github.com/gin-gonic/gin"
)

type OrigStorage struct {
	mx    sync.Mutex
	links map[string]string
}

type ShortenStorage struct {
	mx    sync.Mutex
	links map[string]string
}

var OrigLinks = OrigStorage{links: map[string]string{}}
var ShortenLinks = ShortenStorage{links: map[string]string{}}

// Генерация случайного сокращённого идентификатора
func generateShort() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	short := make([]byte, 8)
	for i := range short {
		short[i] = charset[rand.Intn(len(charset))]
	}
	return string(short)
}

func postLink(c *gin.Context) {
	u, err := c.GetRawData()
	url := string(u)
	if err != nil {
		c.String(http.StatusBadRequest, "Некорректный запрос")
		return
	}

	short := generateShort()
	shortURL := "http://localhost:8080/" + short
	if val, ok := OrigLinks.links[url]; ok {
		c.String(http.StatusCreated, val)
		return
	}

	OrigLinks.mx.Lock()
	OrigLinks.links[url] = shortURL
	OrigLinks.mx.Unlock()

	ShortenLinks.mx.Lock()
	ShortenLinks.links[shortURL] = url
	ShortenLinks.mx.Unlock()

	fmt.Println(url)
	fmt.Println(shortURL)
	fmt.Println()

	c.String(http.StatusCreated, shortURL)
}

func getLink(c *gin.Context) {
	u := c.Param("id")
	shortURL := "http://localhost:8080/" + u

	origURL, ok := ShortenLinks.links[shortURL]
	if !ok {
		c.String(http.StatusBadRequest, "Некорректный запрос")
		return
	}

	c.Header("Location", origURL)
	c.String(http.StatusTemporaryRedirect, "")
}

func RouteNotFound(c *gin.Context) {
	c.String(http.StatusBadRequest, "Некорректный запрос")
}

func main() {
	router := gin.Default()
	router.POST("/", postLink)
	router.GET(":id", getLink)
	router.NoRoute(RouteNotFound)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
