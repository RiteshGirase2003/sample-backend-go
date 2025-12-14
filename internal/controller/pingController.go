// internal/controller/pingController.go
package controller

import (
	"account-guru/internal/response"
	"account-guru/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	var resp response.Response

	if err := service.PingService(); err != nil {
		resp.ResponseWithError(
			response.PING,
			response.InternalServerError,
			"database not connected",
		)

		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.ResponseWithData(
		response.PING,
		response.OK,
		map[string]string{
			"message": "pong",
		},
	)
	c.JSON(http.StatusOK, resp)
}
