package models

import "gorm.io/gorm"

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(
		&PhoneNumber{},
		&FavoriteNumber{},
	)
	return err
}
