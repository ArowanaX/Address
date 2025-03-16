package delivery

import (
	"Address/internal/domain"
	"Address/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CityRegionHandler struct {
	usecase *usecase.CityRegionUseCase
}

func NewCityRegionHandler(e *echo.Echo, uc *usecase.CityRegionUseCase) {
	handler := &CityRegionHandler{usecase: uc}
	e.POST("/api/v1/region", handler.CreateRegion)
}

func (h *CityRegionHandler) CreateRegion(c echo.Context) error {
	var region domain.CityRegion
	if err := c.Bind(&region); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	err := h.usecase.CreateRegion(&region)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create region")
	}

	return c.JSON(http.StatusCreated, region)
}
