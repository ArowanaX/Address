package repo

import (
	"Address/config"
	"Address/internal/domain"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// migrations
	if err := db.AutoMigrate(&domain.CityRegion{}, &domain.MemberAddress{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	return db
}
