package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
)

type Alert struct {
    Labels      map[string]string `json:"labels"`
    Annotations map[string]string `json:annotations`
    StartsAt    time.Time         `json:"startsAt"`
    EndsAt      time.Time         `json:"endsAt"`
}

type Notification struct {
    Version           string            `json:"version"`
    GroupKey          string            `json:"groupKey"`
    Status            string            `json:"status"`
    Receiver          string            `json:receiver`
    GroupLabels       map[string]string `json:groupLabels`
    CommonLabels      map[string]string `json:commonLabels`
    CommonAnnotations map[string]string `json:commonAnnotations`
    ExternalURL       string            `json:externalURL`
    Alerts            []Alert           `json:alerts`
}

func main() {
    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
        })
    })
    router.POST("/webhook", func(c *gin.Context) {
        var notification Notification
        err := c.BindJSON(&notification)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": " successful receive alert notification message!"})
    })
    router.POST("/", func(c *gin.Context) {
        var notification Notification
        err := c.BindJSON(&notification)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": " successful receive alert notification message!"})
    })
    router.Run(":5001")
}

