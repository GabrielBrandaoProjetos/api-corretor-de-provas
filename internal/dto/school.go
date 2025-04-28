package dto

import (
	"time"

	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/domain"
	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateSchoolRequest struct {
	Name     string `json:"name"`
	Register string `json:"register"`
	Unit     string `json:"unit"`
	Address  string `json:"address"`
}

type SchoolResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Register  string    `json:"register"`
	Unit      string    `json:"unit"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToSchool(request CreateSchoolRequest) *domain.School {
	return domain.NewSchool(
		request.Name,
		request.Register,
		request.Unit,
		request.Address,
	)
}

func ToSqlcParams(school *domain.School) repository.CreateParams {
	return repository.CreateParams{
		ID:        school.ID,
		Name:      school.Name,
		Register:  school.Register,
		Unit:      string(school.Unit),
		Address:   school.Address,
		CreatedAt: pgtype.Timestamp{Time: school.CreatedAt, Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: school.UpdatedAt, Valid: true},
	}
}

func FromSchool(school *domain.School) *SchoolResponse {
	return &SchoolResponse{
		ID:        school.ID,
		Name:      school.Name,
		Register:  school.Register,
		Unit:      string(school.Unit),
		Address:   school.Address,
		CreatedAt: school.CreatedAt,
		UpdatedAt: school.UpdatedAt,
	}
}
