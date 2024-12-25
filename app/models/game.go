package models

import (
	"database/sql"
	"time"

	"gorm.io/datatypes"
)

// Card is a generic model that can represent a card from any TCG
type Card struct {
	ID     string `gorm:"primaryKey"`
	Name   string `gorm:"index"`
	Rarity string `gorm:"index"`
	Game   string `gorm:"index"`

	ExtraData datatypes.JSON `gorm:"type:jsonb"`

	// GORM Timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
