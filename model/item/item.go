package item

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint32    `json:"id" gorm:"primarykey"`
	Name        string    `json:"name" gorm:"type:varchar(64)"`
	Price       int       `json:"price"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Item{})
}
