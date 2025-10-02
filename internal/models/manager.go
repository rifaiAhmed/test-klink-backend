package models

import (
	"time"
)

type Manager struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama       string `json:"nama" gorm:"type:varchar(50);not null"`
	LocationID uint   `json:"location_id" gorm:"not null"`

	Location Location `gorm:"foreignKey:LocationID;references:ID" json:"location,omitempty"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Manager) TableName() string {
	return "managers"
}
