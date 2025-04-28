package main

import (
	"context"
	"log"

	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/configs"
	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/service"
	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/util/stringutil"
	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/web/server"
	_ "github.com/lib/pq"
)

func main() {
	configs.InitDB()

	repository := configs.GetRepository()
	schoolService := service.NewSchoolService(context.Background(), repository)

	port := stringutil.GetEnv("HTTP_PORT", "8080")
	svr := server.NewServer(port, schoolService)
	if err := svr.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
