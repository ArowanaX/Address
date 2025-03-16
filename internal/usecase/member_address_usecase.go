package usecase

import (
	"Address/internal/domain"
	"Address/internal/repo"
	"errors"
)

type MemberAddressUseCase struct {
	repo *repo.MemberAddressRepo
}

func NewMemberAddressUseCase(r *repo.MemberAddressRepo) *MemberAddressUseCase {
	return &MemberAddressUseCase{repo: r}
}

// Create a new address
func (u *MemberAddressUseCase) CreateAddress(address *domain.MemberAddress) error {
	return u.repo.Create(address)
}

// Edit an address
func (u *MemberAddressUseCase) EditAddress(address *domain.MemberAddress) error {
	return u.repo.Update(address)
}

// Select an address (ensuring only one is selected)
func (u *MemberAddressUseCase) SelectAddress(addressID uint) error {
	address, err := u.repo.GetByID(addressID)
	if err != nil {
		return errors.New("address not found")
	}

	// Unselect all addresses for the user
	if err := u.repo.UnselectAllForMember(address.Member); err != nil {
		return err
	}

	// Select this address
	address.IsSelected = true
	return u.repo.Update(address)
}
