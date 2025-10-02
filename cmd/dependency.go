package cmd

import (
	"backend-test/helpers"
	"backend-test/internal/api"
	"backend-test/internal/interfaces"
	repository "backend-test/internal/repositories"
	service "backend-test/internal/services"
)

type Dependency struct {
	MemberAPI interfaces.IMemberHandler
}

func dependencyInject() Dependency {
	memberRepo := &repository.MemberRepository{
		DB: helpers.DB,
	}

	memberSvc := service.NewMemberService(memberRepo)

	memberAPI := &api.IMemberHandler{
		MemberService: memberSvc,
	}

	return Dependency{
		MemberAPI: memberAPI,
	}
}
