package v1

import (
	"kode-notes/internal/service"
	"time"
)

type Handler struct {
	services *service.Service
	tokenTTL time.Duration
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
