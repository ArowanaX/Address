package domain

import "gorm.io/gorm"

type CityRegion struct {
	gorm.Model
	ID     uint         `gorm:"primaryKey"`
	Number int          `gorm:"not null"`
	Parent *uint        `gorm:"default:null"`
	Name   string       `gorm:"size:255;not null"`
	Childs []CityRegion `gorm:"foreignKey:Parent"`
}

type MemberAddress struct {
	gorm.Model
	ID            uint       `gorm:"primaryKey"`
	Member        int        `gorm:"not null"`
	PhoneNumber   string     `gorm:"size:20;not null"`
	IsSelected    bool       `gorm:"default:false"`
	PostalAddress string     `gorm:"type:text;not null"`
	RegionID      uint       `gorm:"not null"`
	Region        CityRegion `gorm:"foreignKey:RegionID"`
}
