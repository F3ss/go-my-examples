package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) signUp(c *gin.Context) {
	c.String(http.StatusOK, "Ok")
}

func (h *handler) signIn(c *gin.Context) {
	c.String(http.StatusOK, "Ok")
}
