package services

import (
	"net/http"
	"strconv"
	"tcg-games/app/db"

	"github.com/labstack/echo/v4"
)

// ListCards handles GET /cards?name=...&rarities=...&games=...&colors=...&ink_cost_min=...&ink_cost_max=...
func (gs *gameService) ListCards(c echo.Context) error {
	// Build up our filter struct
	filters := db.CardFilters{}

	// 1. Partial name match (if provided)
	if name := c.QueryParam("name"); name != "" {
		filters.Name = &name
	}

	rarityParams := c.QueryParams()["rarities"]
	if len(rarityParams) > 0 {
		filters.Rarities = rarityParams
	}

	gamesParams := c.QueryParams()["games"]
	if len(gamesParams) > 0 {
		filters.Games = gamesParams
	}

	colorsParams := c.QueryParams()["colors"]
	if len(colorsParams) > 0 {
		filters.Colors = colorsParams
	}

	if inkCostMinStr := c.QueryParam("ink_cost_min"); inkCostMinStr != "" {
		if val, err := strconv.Atoi(inkCostMinStr); err == nil {
			filters.InkCostMin = &val
		}
	}
	if inkCostMaxStr := c.QueryParam("ink_cost_max"); inkCostMaxStr != "" {
		if val, err := strconv.Atoi(inkCostMaxStr); err == nil {
			filters.InkCostMax = &val
		}
	}

	cards, err := gs.store.ListCards(filters)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cards)

}
