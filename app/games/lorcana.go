package games

import "tcg-games/app/enums"

type LorcanaCard struct {
	BaseGame
	Rarity  enums.LorcanaRarity `json:"rarity"`
	InkCost int                 `json:"ink_cost"`
}

func (l *LorcanaCard) GetRarity() string {
	return string(l.Rarity)
}
