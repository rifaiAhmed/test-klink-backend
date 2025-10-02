package service

import (
	"backend-test/internal/interfaces"
	"backend-test/internal/models"
	"context"
	"errors"
)

type MemberService struct {
	Repo interfaces.IMemberRepository
}

func NewMemberService(repo interfaces.IMemberRepository) *MemberService {
	return &MemberService{
		Repo: repo,
	}
}

func (s *MemberService) CreateMember(ctx context.Context, member *models.Member, paketID uint) (*models.Member, error) {
	if err := member.Validate(); err != nil {
		return nil, err
	}

	saved, err := s.Repo.Save(ctx, member, paketID)
	if err != nil {
		return nil, err
	}

	return saved, nil
}

func (s *MemberService) GetMemberByID(ctx context.Context, id string) (*models.Member, error) {
	member, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("member tidak ditemukan")
	}
	return member, nil
}

func (s *MemberService) UpdateMember(ctx context.Context, user *models.Member) (*models.Member, error) {
	existing, err := s.Repo.FindByID(ctx, user.ID)
	if err != nil {
		return nil, errors.New("member tidak ditemukan")
	}

	existing.Nama = user.Nama
	existing.JenisKelamin = user.JenisKelamin
	existing.NoKtp = user.NoKtp
	existing.TempatLahir = user.TempatLahir
	existing.TanggalLahir = user.TanggalLahir
	existing.NoHp = user.NoHp
	existing.Email = user.Email
	existing.NoRekening = user.NoRekening
	existing.ManagerID = user.ManagerID

	if err := existing.Validate(); err != nil {
		return nil, err
	}

	return s.Repo.Save(ctx, existing, existing.Registration.PaketID)
}

func (s *MemberService) DeleteMember(ctx context.Context, id string) error {
	return s.Repo.DeleteMember(ctx, id)
}

func (s *MemberService) GetAllMembers(ctx context.Context, objComponent models.ComponentServerSide) ([]models.Member, int64, error) {
	members, err := s.Repo.GetAllMembers(ctx, objComponent, "")
	if err != nil {
		return nil, 0, err
	}

	count, err := s.Repo.CountData(ctx, objComponent)
	if err != nil {
		return nil, 0, err
	}

	return members, count, nil
}
func (s *MemberService) GetManagers() ([]models.Option, error) {
	data, err := s.Repo.GetManagers()
	return data, err
}

func (s *MemberService) GetPakets() ([]models.Option, error) {
	data, err := s.Repo.GetPakets()
	return data, err
}

func (s *MemberService) GetMembers() ([]models.Option2, error) {
	data, err := s.Repo.GetMembers()
	return data, err
}
