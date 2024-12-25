package main

import (
	"tcg-games/app/db/postgres"
	"tcg-games/app/helpers"
	"tcg-games/app/routes"
	"tcg-games/app/seed"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("Error loading environment  variables file!")
	}

	postgres, err := postgres.New()
	if err != nil {
		log.Fatal().Err(err).Msg("error connecting to db")
	}

	// seeding the data:
	err = seed.SeedCards(postgres, "lorcana-cards.json", "mtg-cards.json")
	if err != nil {
		log.Error().Err(err).Msg("error seeding cards, if you have already seeded the data, you can ignore this error")
	}
	// end seeding

	port := helpers.Env("PORT")
	if port == "" {
		port = "9090"
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize routes
	routes.InitRoutes(e, postgres)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
