package main

import (
	"fmt"
	"time"

	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

// redisFunc does insane shiz
func redisAdd(url string) string {

	randid := randSeq(6)
	// if val == "" {
	err := client.Set(randid, url, time.Hour/2).Err()
	if err != nil {
		panic(err)
	}
	return randid
}

func redisGet(shorturl string) string {
	val, err := client.Get(shorturl).Result()
	if err != nil {
		fmt.Println("Error")
		// panic(err)
	}
	return val
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/add/:url", func(c *gin.Context) {
		url := c.Params.ByName("url")
		retURL := redisAdd(url)
		c.String(200, retURL)
	})

	r.GET("/api/get/:shorturl", func(c *gin.Context) {
		shorturl := c.Params.ByName("shorturl")
		retURL := redisGet(shorturl)
		c.String(200, retURL)
	})
	r.Run()
}
