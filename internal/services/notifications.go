package services

import "log"

type NotificationService interface {
	SendNotification()
}

type notificationService struct{}

func NewNotificationService() NotificationService {
	return &notificationService{}
}

func (s *notificationService) SendNotification() {
	log.Println("Processing notification...")
}