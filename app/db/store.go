package db

import "tcg-games/app/models"

type CardFilters struct {
	// If provided, do a partial (case-insensitive) match on the name.
	Name *string

	// If provided, only return cards whose rarity is in this slice.
	Rarities []string

	// If provided, only return cards whose game is in this slice.
	Games []string

	// Example of an additional filter stored in ExtraData:
	// If provided, only return cards whose "color" attribute
	// (in ExtraData->>'color') is in this slice.
	Colors []string

	// For numeric range filtering on ink_cost (in ExtraData->>'ink_cost').
	// If provided, only return cards with ink_cost >= InkCostMin.
	InkCostMin *int
	// If provided, only return cards with ink_cost <= InkCostMax.
	InkCostMax *int
}

type Store interface {
	ListCards(filters CardFilters) ([]models.Card, error)
	AddCards(cards []models.Card) error
}
