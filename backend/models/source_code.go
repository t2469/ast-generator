package models

import "time"

type SourceCode struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description" gorm:"default:''"`
	Language    string    `json:"language"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"created_at"`
}

func (SourceCode) TableName() string {
	return "source_codes"
}
