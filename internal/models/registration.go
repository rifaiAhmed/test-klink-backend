package models

import (
	"time"
)

type Registration struct {
	ID             uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	MemberID       string `json:"member_id" gorm:"type:varchar(50);not null"`
	UplineMemberID string `json:"upline_member_id" gorm:"type:varchar(50);"`
	PaketID        uint   `json:"paket_id" gorm:"not null"`

	Paket Paket `gorm:"foreignKey:PaketID;references:ID" json:"paket,omitempty"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Registration) TableName() string {
	return "registrations"
}
