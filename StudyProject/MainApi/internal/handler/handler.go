package handler

import "github.com/gin-gonic/gin"

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/a", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	auth := router.Group("api/v1/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	return router
}
