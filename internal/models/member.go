package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Member struct {
	ID           string `json:"id" gorm:"primaryKey;type:varchar(50)"`
	Nama         string `json:"nama" gorm:"type:varchar(50);not null" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" gorm:"type:varchar(20);not null" validate:"required"`
	NoKtp        string `json:"no_ktp" gorm:"type:varchar(50);unique;not null" validate:"required"`
	TempatLahir  string `json:"tempat_lahir" gorm:"type:varchar(50);not null" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" gorm:"type:varchar(10);not null" validate:"required"`
	NoHp         string `json:"no_hp" gorm:"type:varchar(20);"`
	Email        string `json:"email" gorm:"type:varchar(50);"`
	NoRekening   string `json:"no_rekening" gorm:"type:varchar(50);"`

	ManagerID uint    `json:"manager_id" gorm:"not null"`
	Manager   Manager `gorm:"foreignKey:ManagerID;references:ID" json:"manager,omitempty"`

	RegistartionID uint         `json:"Registartion_id" gorm:"not null"`
	Registration   Registration `gorm:"foreignKey:MemberID;references:ID" json:"registration,omitempty"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (Member) TableName() string {
	return "members"
}

func (m Member) Validate() error {
	v := validator.New()
	return v.Struct(m)
}
