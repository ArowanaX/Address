package delivery

import (
	"Address/internal/domain"
	"Address/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MemberAddressHandler struct {
	usecase *usecase.MemberAddressUseCase
}

func NewMemberAddressHandler(e *echo.Echo, uc *usecase.MemberAddressUseCase) {
	handler := &MemberAddressHandler{usecase: uc}
	e.POST("/api/v1/address", handler.CreateAddress)
	e.PUT("/api/v1/address/:id", handler.EditAddress)
	e.PATCH("/api/v1/address/:id/select", handler.SelectAddress)
}

// Create a new address
func (h *MemberAddressHandler) CreateAddress(c echo.Context) error {
	var address domain.MemberAddress
	if err := c.Bind(&address); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	err := h.usecase.CreateAddress(&address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create address")
	}

	return c.JSON(http.StatusCreated, address)
}

// Edit an address
func (h *MemberAddressHandler) EditAddress(c echo.Context) error {
	var address domain.MemberAddress
	if err := c.Bind(&address); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	err := h.usecase.EditAddress(&address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to edit address")
	}

	return c.JSON(http.StatusOK, address)
}

// Select an address (ensuring only one is selected)
func (h *MemberAddressHandler) SelectAddress(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid address ID")
	}

	err = h.usecase.SelectAddress(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to select address")
	}

	return c.JSON(http.StatusOK, "Address selected successfully")
}
