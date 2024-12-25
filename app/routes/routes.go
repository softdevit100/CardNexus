package routes

import (
	"tcg-games/app/db"
	"tcg-games/app/services"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, store db.Store) {

	gameService := services.NewGameService(store)

	e.GET("/cards", gameService.ListCards)
}
