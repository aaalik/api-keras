package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint32    `json:"id" gorm:"primarykey"`
	Name      string    `json:"name" gorm:"type:varchar(64)"`
	Email     string    `json:"email" gorm:"type:varchar(64)"`
	Password  string    `json:"password" gorm:"type:varchar(64)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
