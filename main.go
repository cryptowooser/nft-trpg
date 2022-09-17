package main

import (
	"net/http"
	"time"

	"example.com/nft-trpg/stats"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	r.POST("/post", func(ctx *gin.Context) {

		type postResponse struct {
			Time         int64  `json:"time"`
			Message      string `json:"message"`
			HelloMessage string `json:"hello_message"`
		}

		resp := postResponse{
			Time:         time.Now().Unix(),
			Message:      "Test message",
			HelloMessage: "Hello message!",
		}

		ctx.JSON(http.StatusOK, resp)
	})

	r.POST("/get-stats", func(ctx *gin.Context) {

		type statRequest struct {
			ID int64 `json:"id"`
		}

		var req statRequest

		err := ctx.ShouldBind(&req)
		if err != nil {
			panic(err)
		}

		stats := stats.GetStats(req.ID)

		ctx.JSON(http.StatusOK, stats)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
