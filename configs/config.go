package configs

import "github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/repository"

var db *repository.Queries

func InitDB() {
	db = InitPostgres()
}

func GetRepository() *repository.Queries {
	return db
}
