package services

import (
	"tcg-games/app/db"

	"github.com/labstack/echo/v4"
)

type Service interface {
	ListCards(c echo.Context) error
}

type gameService struct {
	store db.Store
}

func NewGameService(
	store db.Store,
) Service {
	return &gameService{
		store: store,
	}
}
