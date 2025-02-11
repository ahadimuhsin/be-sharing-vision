package post

import (
	"database/sql/driver"
	"time"
)

type status string

const (
	Publish status = "publish"
	Draft   status = "draft"
	Trash   status = "trash"
)

func (st *status) Scan(value interface{}) error {
	*st = status(value.([]byte))
	return nil
}

func (st status) Value() (driver.Value, error) {
	return string(st), nil
}

type Post struct {
	ID        int64     `json:"id"`
	Title     string    `gorm:"type:varchar(200)" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Category  string    `gorm:"type:varchar(100)" json:"category"`
	CreatedAt time.Time `gorm:"type:datetime;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;default:current_timestamp" json:"updated_at"`
	Status    status    `gorm:"type:enum('publish', 'draft', 'trash')" json:"status"`
}
