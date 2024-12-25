package games

type Game interface {
	GetID() string
	GetName() string
	GetRarity() string
}

type BaseGame struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (bg *BaseGame) GetID() string {
	return bg.ID
}

func (bg *BaseGame) GetName() string {
	return bg.Name
}
