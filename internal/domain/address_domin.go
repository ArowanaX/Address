package domain

import "gorm.io/gorm"

type CityRegion struct {
	gorm.Model
	Number int          `gorm:"not null;uniqueIndex" validate:"required,min=1"`
	Parent *uint        `gorm:"default:null"`
	Name   string       `gorm:"size:255;not null" validate:"required,min=2,max=255"`
	Childs []CityRegion `gorm:"foreignKey:Parent"`
}

type MemberAddress struct {
	gorm.Model
	Member        int        `gorm:"not null" validate:"required,min=1"`
	PhoneNumber   string     `gorm:"size:20;not null" validate:"required,e164"`
	IsSelected    bool       `gorm:"default:false"`
	PostalAddress string     `gorm:"type:text;not null" validate:"required,min=10"`
	RegionID      uint       `gorm:"not null" validate:"required,min=1"`
	Region        CityRegion `gorm:"foreignKey:RegionID"`
}
