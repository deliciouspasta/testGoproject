package main

import "github.com/gin-gonic/gin"

import "net/http"

func main() {
    router:= gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
    })
    router.Run(":8888")
}