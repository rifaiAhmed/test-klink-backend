package models

import (
	"time"
)

type Paket struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaPaket  string    `json:"nama_paket" gorm:"type:varchar(50);not null"`
	JenisPaket string    `json:"jenis_paket" gorm:"type:varchar(20);not null"`
	Wilayah    string    `json:"wilayah" gorm:"type:varchar(25);not null"`
	Price      float64   `json:"price" gorm:"type:numeric(12,2);not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Paket) TableName() string {
	return "pakets"
}
