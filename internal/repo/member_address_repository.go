package repo

import (
	"Address/internal/domain"
	"gorm.io/gorm"
)

type MemberAddressRepo struct {
	DB *gorm.DB
}

func NewMemberAddressRepo(db *gorm.DB) *MemberAddressRepo {
	return &MemberAddressRepo{DB: db}
}

// Create a new address
func (r *MemberAddressRepo) Create(address *domain.MemberAddress) error {
	return r.DB.Create(address).Error
}

// Update an existing address
func (r *MemberAddressRepo) Update(address *domain.MemberAddress) error {
	return r.DB.Save(address).Error
}

// Select an address by ID
func (r *MemberAddressRepo) GetByID(id uint) (*domain.MemberAddress, error) {
	var address domain.MemberAddress
	err := r.DB.First(&address, id).Error
	return &address, err
}

// Ensure only one selected address per user
func (r *MemberAddressRepo) UnselectAllForMember(memberID int) error {
	return r.DB.Model(&domain.MemberAddress{}).Where("member = ?", memberID).Update("is_selected", false).Error
}
