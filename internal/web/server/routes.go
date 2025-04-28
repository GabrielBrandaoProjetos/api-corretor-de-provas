package server

import (
	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/web/handler"
)

func InitializeRoutes(s *Server) {
	router := s.server
	schoolHandle := handler.NewSchoolHandler(s.schoolService)
	basePath := "/api"

	api := router.Group(basePath)
	{
		api.POST("/school", schoolHandle.CreateSchoolHandler)
	}
}
