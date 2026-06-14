package utils

import (
	"github.com/gin-gonic/gin"
)

type Success struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Error struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SuccessResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Success{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, status int, err string) {
	c.JSON(status, Error{
		Success: false,
		Message:   err,
	})
}