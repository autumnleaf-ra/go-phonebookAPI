package models

type PhoneNumber struct {
	Id        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    string `json:"phone_number"`
}

type FavoriteNumber struct {
	Id          uint        `gorm:"primaryKey;autoIncrement" json:"favorites_id"`
	PhoneID     uint        `json:"phoneid"`
	PhoneNumber PhoneNumber `gorm:"foreignKey:PhoneID"`
}
