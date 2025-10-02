package dto

import "github.com/go-playground/validator/v10"

type CreateMemberRequest struct {
	Nama         string `json:"nama" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" validate:"required"`
	NoKtp        string `json:"no_ktp" validate:"required"`
	TempatLahir  string `json:"tempat_lahir" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	NoHp         string `json:"no_hp"`
	Email        string `json:"email"`
	NoRekening   string `json:"no_rekening"`
	ManagerID    uint   `json:"manager_id" validate:"required"`
	PaketID      uint   `json:"paket_id" validate:"required"`
}

func (r CreateMemberRequest) Validate() error {
	v := validator.New()
	return v.Struct(r)
}
