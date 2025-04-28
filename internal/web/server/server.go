package server

import (
	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	server        *gin.Engine
	schoolService *service.SchoolService
	port          string
}

func NewServer(port string, schoolService *service.SchoolService) *Server {
	return &Server{
		server:        gin.Default(),
		schoolService: schoolService,
		port:          port,
	}
}

func (s *Server) Start() error {
	InitializeRoutes(s)
	return s.server.Run(":" + s.port)
}
