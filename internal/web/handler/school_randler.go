package handler

import (
	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/dto"
	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/service"
	"github.com/gin-gonic/gin"
)

type SchoolHandler struct {
	schoolService *service.SchoolService
}

func NewSchoolHandler(schoolService *service.SchoolService) *SchoolHandler {
	return &SchoolHandler{schoolService: schoolService}
}

func (h *SchoolHandler) CreateSchoolHandler(ctx *gin.Context) {
	request := dto.CreateSchoolRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	school, err := h.schoolService.Create(request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"data": school})
}
