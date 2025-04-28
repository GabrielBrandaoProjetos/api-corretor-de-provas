package service

import (
	"context"
	"fmt"

	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/domain"
	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/dto"
	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/repository"
)

type SchoolService struct {
	ctx        context.Context
	repository *repository.Queries
}

func NewSchoolService(ctx context.Context, repository *repository.Queries) *SchoolService {
	return &SchoolService{
		repository: repository,
		ctx:        ctx,
	}
}

func (s *SchoolService) Create(request dto.CreateSchoolRequest) (*dto.SchoolResponse, error) {
	school := dto.ToSchool(request)

	existingSchool, err := s.repository.FindByRegister(s.ctx, school.Register)

	if err != nil && err.Error() != domain.ErrSchoolNotFound.Error() {
		return nil, fmt.Errorf("error while searching for existing school: %v", err)
	}

	if (existingSchool != repository.School{}) && existingSchool.Unit == domain.UnitHead {
		return nil, domain.ErrDuplicatedSchool
	}

	schoolParams := dto.ToSqlcParams(school)

	if err := s.repository.Create(s.ctx, schoolParams); err != nil {
		return nil, fmt.Errorf("error while creating school: %v", err)
	}

	return dto.FromSchool(school), nil
}
