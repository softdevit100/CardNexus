package postgres

import "tcg-games/app/models"

func (p *Postgres) AddCards(cards []models.Card) error {
	// Start a transaction
	tx := p.db.Begin()

	// Insert each card into the "cards" table
	for _, card := range cards {
		if err := tx.Create(&card).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	return tx.Commit().Error
}
