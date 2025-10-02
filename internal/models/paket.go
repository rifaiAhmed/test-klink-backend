package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Paket struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaPaket  string    `json:"nama_paket" gorm:"type:varchar(50);not null" validate:"required"`
	JenisPaket string    `json:"jenis_paket" gorm:"type:varchar(20);not null" validate:"required"`
	Wilayah    string    `json:"wilayah" gorm:"type:varchar(25);not null" validate:"required"`
	Price      float64   `json:"price" gorm:"type:numeric(12,2);not null" validate:"required,gt=0"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Paket) TableName() string {
	return "pakets"
}

func (p Paket) Validate() error {
	v := validator.New()
	return v.Struct(p)
}
