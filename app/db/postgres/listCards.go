package postgres

import (
	"tcg-games/app/db"
	"tcg-games/app/models"
)

func (p *Postgres) ListCards(filters db.CardFilters) ([]models.Card, error) {
	query := p.db.Model(&models.Card{})

	if filters.Name != nil && *filters.Name != "" {
		query = query.Where("name ILIKE ?", "%"+*filters.Name+"%")
	}

	if len(filters.Rarities) > 0 {
		query = query.Where("rarity IN ?", filters.Rarities)
	}

	if len(filters.Games) > 0 {
		query = query.Where("game IN ?", filters.Games)
	}

	if len(filters.Colors) > 0 {
		query = query.Where("(extra_data->>'color') IN ?", filters.Colors)
	}

	if filters.InkCostMin != nil {
		query = query.Where("(extra_data->>'ink_cost')::int >= ?", *filters.InkCostMin)
	}
	if filters.InkCostMax != nil {
		query = query.Where("(extra_data->>'ink_cost')::int <= ?", *filters.InkCostMax)
	}

	var cards []models.Card
	if err := query.Find(&cards).Error; err != nil {
		return nil, err
	}

	return cards, nil
}
