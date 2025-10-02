package dto

import (
	"github.com/go-playground/validator/v10"
)

type UpdateMemberRequest struct {
	Nama         string `json:"nama" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" validate:"required"`
	NoKtp        string `json:"no_ktp" validate:"required"`
	TempatLahir  string `json:"tempat_lahir" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	NoHp         string `json:"no_hp"`
	Email        string `json:"email"`
	NoRekening   string `json:"no_rekening"`

	ManagerID      uint `json:"manager_id" validate:"required"`
	RegistrationID uint `json:"registration_id" validate:"required"`
}

func (r *UpdateMemberRequest) Validate() error {
	v := validator.New()
	return v.Struct(r)
}
