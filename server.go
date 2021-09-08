package main

import (
	"demo/controller"
	"demo/middlewares"
	"demo/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger()) // TODO: change logging format

	heathCheck := server.Group("/health")
	{
		heathCheck.GET("/ping", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "pong")
		})
	}
	apiRoutes := server.Group("/api", middlewares.BasicAuth())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video Input is Valid!",
				})
			}

		})
	}

	// We can set up this env variable from the Elastic Beanstalk Console
	port := os.Getenv("PORT")
	// EB forwards requests to port 5000
	if port == "" {
		port = "5000"
	}


	server.Run(":" + port)
}
