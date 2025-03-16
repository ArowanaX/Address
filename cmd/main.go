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
	// Initialize repositories
	regionRepo := repo.NewCityRegionRepo(db)
	addressRepo := repo.NewMemberAddressRepo(db)

	// Initialize use cases
	regionUC := usecase.NewCityRegionUseCase(regionRepo)
	addressUC := usecase.NewMemberAddressUseCase(addressRepo)

	// Register handlers
	delivery.NewCityRegionHandler(e, regionUC)
	delivery.NewMemberAddressHandler(e, addressUC)

	err := e.Start(":8000")
	if err != nil {
		return
	}
}
