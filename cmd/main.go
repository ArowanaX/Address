package main

import (
	"Address/config"
	"Address/internal/delivery"
	"Address/internal/repo"
	"Address/internal/usecase"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()
	db := repo.InitDB(cfg)
	fmt.Println("Database connected!")

	e := echo.New()
	regionRepo := repo.NewCityRegionRepo(db)
	regionUC := usecase.NewCityRegionUseCase(regionRepo)
	delivery.NewCityRegionHandler(e, regionUC)

	err := e.Start(":8000")
	if err != nil {
		return
	}
}
