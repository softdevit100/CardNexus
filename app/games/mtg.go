package games

import "tcg-games/app/enums"

// MTGRarity is an enum-like type for Magic: The Gathering rarities.

// MTGCard represents a Magic: The Gathering card.
type MTGCard struct {
	BaseGame
	Rarity enums.MTGRarity `json:"rarity"`
	Color  *string         `json:"color,omitempty"` // e.g. "B", "R", "U", "W", "G"
}

// GetRarity returns the rarity as a string (to satisfy the Game interface).
func (m *MTGCard) GetRarity() string {
	return string(m.Rarity)
}
