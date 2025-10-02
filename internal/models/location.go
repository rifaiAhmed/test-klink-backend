package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Location struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Kelurahan string    `json:"kelurahan" gorm:"type:varchar(50);not null" validate:"required"`
	Kecamatan string    `json:"kecamatan" gorm:"type:varchar(50);not null" validate:"required"`
	Kabupaten string    `json:"kabupaten" gorm:"type:varchar(50);not null" validate:"required"`
	KodePos   string    `json:"kode_pos" gorm:"type:varchar(10);not null" validate:"required"`
	Detail    string    `json:"detail" gorm:"type:varchar(100);"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Location) TableName() string {
	return "locations"
}

func (l Location) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
