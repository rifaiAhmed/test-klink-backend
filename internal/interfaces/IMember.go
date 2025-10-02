package interfaces

import (
	"backend-test/internal/models"
	"context"

	"github.com/gin-gonic/gin"
)

type IMemberRepository interface {
	Save(ctx context.Context, user *models.Member) (*models.Member, error)
	FindByID(ctx context.Context, id string) (*models.Member, error)
	DeleteMember(ctx context.Context, id string) error
	GetAllMembers(ctx context.Context, objComponent models.ComponentServerSide, param string) ([]models.Member, error)
	CountData(ctx context.Context, objComponent models.ComponentServerSide) (int64, error)
	CreateRegistration(ctx context.Context, reg *models.Registration) error
}

type IMemberService interface {
	CreateMember(ctx context.Context, member *models.Member, paketID uint) (*models.Member, error)
	GetMemberByID(ctx context.Context, id string) (*models.Member, error)
	DeleteMember(ctx context.Context, id string) error
	UpdateMember(ctx context.Context, user *models.Member) (*models.Member, error)
	GetAllMembers(ctx context.Context, objComponent models.ComponentServerSide) ([]models.Member, int64, error)
}

type IMemberHandler interface {
	CreateMember(c *gin.Context)
	UpdateMember(c *gin.Context)
	GetMemberByID(c *gin.Context)
	DeleteMember(c *gin.Context)
	GetAllMembers(c *gin.Context)
}
