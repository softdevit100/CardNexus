package seed

import (
	"encoding/json"
	"fmt"
	"os"
	"tcg-games/app/db"
	"tcg-games/app/models"
)

// SeedCards reads two JSON files (Lorcana + MTG), converts them into []models.Card,
func SeedCards(store db.Store, lorcanaPath, mtgPath string) error {
	// 1) Seed Lorcana Cards
	lorcanaCards, err := readLorcanaCards(lorcanaPath)
	if err != nil {
		return fmt.Errorf("failed to read Lorcana cards: %w", err)
	}
	if err := store.AddCards(lorcanaCards); err != nil {
		return fmt.Errorf("failed to add Lorcana cards: %w", err)
	}

	// 2) Seed MTG Cards
	mtgCards, err := readMTGCards(mtgPath)
	if err != nil {
		return fmt.Errorf("failed to read MTG cards: %w", err)
	}
	if err := store.AddCards(mtgCards); err != nil {
		return fmt.Errorf("failed to add MTG cards: %w", err)
	}

	return nil
}

func readLorcanaCards(path string) ([]models.Card, error) {
	type lorcanaCard struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Rarity  string `json:"rarity"`
		InkCost int    `json:"ink_cost"`
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var rawCards []lorcanaCard
	if err := json.Unmarshal(data, &rawCards); err != nil {
		return nil, err
	}

	var results []models.Card
	for _, rc := range rawCards {
		extra := map[string]interface{}{
			// "external_id": rc.ID,
			"ink_cost": rc.InkCost,
		}
		extraBytes, _ := json.Marshal(extra)

		card := models.Card{
			ID:        rc.ID,
			Name:      rc.Name,
			Rarity:    rc.Rarity,
			Game:      "Lorcana",
			ExtraData: extraBytes,
		}
		results = append(results, card)
	}
	return results, nil
}

func readMTGCards(path string) ([]models.Card, error) {
	type mtgCard struct {
		ID     string  `json:"id"`
		Name   string  `json:"name"`
		Rarity string  `json:"rarity"`
		Color  *string `json:"color,omitempty"`
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var rawCards []mtgCard
	if err := json.Unmarshal(data, &rawCards); err != nil {
		return nil, err
	}

	var results []models.Card
	for _, rc := range rawCards {
		extra := map[string]interface{}{
			// "external_id": rc.ID,
		}
		if rc.Color != nil && *rc.Color != "" {
			extra["color"] = *rc.Color
		}
		extraBytes, _ := json.Marshal(extra)

		card := models.Card{
			ID:        rc.ID,
			Name:      rc.Name,
			Rarity:    rc.Rarity,
			Game:      "MTG",
			ExtraData: extraBytes,
		}
		results = append(results, card)
	}
	return results, nil
}
