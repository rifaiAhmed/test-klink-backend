package repository

import (
	"backend-test/internal/models"
	"context"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type MemberRepository struct {
	DB *gorm.DB
}

func (r *MemberRepository) generateMemberID(ctx context.Context) (string, error) {
	var count int64
	today := time.Now().Format("20060102")

	if err := r.DB.WithContext(ctx).Model(&models.Member{}).
		Where("DATE(created_at) = CURRENT_DATE").
		Count(&count).Error; err != nil {
		return "", err
	}

	newID := fmt.Sprintf("MEM-%s-%04d", today, count+1)
	return newID, nil
}

func (r *MemberRepository) Save(ctx context.Context, member *models.Member, paketID uint) (*models.Member, error) {
	if member.ID == "" {
		id, err := r.generateMemberID(ctx)
		if err != nil {
			return nil, err
		}
		member.ID = id
	}

	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(member).Error; err != nil {
			return err
		}

		reg := &models.Registration{
			MemberID: member.ID,
			PaketID:  paketID,
		}
		if err := tx.Create(reg).Error; err != nil {
			return err
		}

		member.RegistartionID = reg.ID
		if err := tx.Save(member).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if err := r.DB.Preload("Manager").Preload("Manager.Location").Preload("Registration.Paket").First(member, "id = ?", member.ID).Error; err != nil {
		return nil, err
	}

	return member, nil
}

func (r *MemberRepository) CreateRegistration(ctx context.Context, reg *models.Registration) error {
	return r.DB.WithContext(ctx).Create(reg).Error
}

func (r *MemberRepository) FindByID(ctx context.Context, id string) (*models.Member, error) {
	var obj models.Member
	if err := r.DB.WithContext(ctx).Preload("Manager").Preload("Manager.Location").Preload("Registration.Paket").Where("id = ?", id).First(&obj).Error; err != nil {
		return nil, err
	}
	return &obj, nil
}

func (r *MemberRepository) DeleteMember(ctx context.Context, id string) error {
	return r.DB.WithContext(ctx).Where("id = ?", id).Delete(&models.Member{}).Error
}

func (r *MemberRepository) GetAllMembers(ctx context.Context, objComponent models.ComponentServerSide, param string) ([]models.Member, error) {
	var members []models.Member

	limit := objComponent.Limit
	offset := objComponent.Skip
	order := fmt.Sprintf("%s %s", objComponent.SortBy, objComponent.SortType)

	db := r.DB.WithContext(ctx).
		Model(&models.Member{}).
		Preload("Manager").
		Preload("Manager.Location").
		Preload("Registration").
		Preload("Registration.Paket")

	if objComponent.Search != "" {
		search := "%" + strings.ToLower(objComponent.Search) + "%"
		db = db.Where("LOWER(nama) LIKE ?", search)
	}

	if err := db.Order(order).Limit(limit).Offset(offset).Find(&members).Error; err != nil {
		return nil, err
	}

	return members, nil
}

func (r *MemberRepository) CountData(ctx context.Context, objComponent models.ComponentServerSide) (int64, error) {
	var count int64

	db := r.DB.WithContext(ctx).Model(&models.Member{})

	if objComponent.Search != "" {
		search := "%" + strings.ToLower(objComponent.Search) + "%"
		db = db.Where("LOWER(nama) LIKE ?", search)
	}

	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (r *MemberRepository) GetManagers() ([]models.Option, error) {
	var (
		managers []models.Manager
		options  []models.Option
	)

	if err := r.DB.Find(&managers).Error; err != nil {
		return nil, err
	}
	for _, m := range managers {
		options = append(options, models.Option{
			Value: int(m.ID),
			Label: m.Nama,
		})
	}

	return options, nil
}

func (r *MemberRepository) GetPakets() ([]models.Option, error) {
	var (
		pakets  []models.Paket
		options []models.Option
	)

	if err := r.DB.Find(&pakets).Error; err != nil {
		return nil, err
	}
	for _, m := range pakets {
		options = append(options, models.Option{
			Value: int(m.ID),
			Label: m.NamaPaket,
		})
	}

	return options, nil
}

func (r *MemberRepository) GetMembers() ([]models.Option2, error) {
	var (
		members []models.Member
		options []models.Option2
	)

	if err := r.DB.Find(&members).Error; err != nil {
		return nil, err
	}
	for _, m := range members {
		options = append(options, models.Option2{
			Value: m.ID,
			Label: m.Nama,
		})
	}

	return options, nil
}
