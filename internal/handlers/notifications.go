package handlers

import (
	"net/http"

	"github.com/Habeebamoo/tunnl-backend/internal/services"
	"github.com/Habeebamoo/tunnl-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	service services.NotificationService
}

func NewNotificationHandler(s services.NotificationService) *NotificationHandler {
	return &NotificationHandler{service: s}
}

func (h *NotificationHandler) SendNotification(c *gin.Context) {
	h.service.SendNotification()
	utils.SuccessResponse(c, http.StatusOK, "Notification sent successfully", nil)
}