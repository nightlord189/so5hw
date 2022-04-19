package handler

import (
	"github.com/nightlord189/so5hw/internal/config"
	database "github.com/nightlord189/so5hw/internal/db"
)

type Handler struct {
	Config *config.Config
	DB     *database.Manager
}

func NewHandler(
	cfg *config.Config,
	db *database.Manager) *Handler {
	handler := Handler{
		Config: cfg,
		DB:     db,
	}
	return &handler
}
