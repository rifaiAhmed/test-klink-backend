package models

import (
	"time"
)

type Location struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Kelurahan string    `json:"kelurahan" gorm:"type:varchar(50);not null"`
	Kecamatan string    `json:"kecamatan" gorm:"type:varchar(50);not null"`
	Kabupaten string    `json:"kabupaten" gorm:"type:varchar(50);not null"`
	KodePos   string    `json:"kode_pos" gorm:"type:varchar(10);not null"`
	Detail    string    `json:"detail" gorm:"type:varchar(100);"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Location) TableName() string {
	return "locations"
}
