package models

import "time"

type SourceCode struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Language  string    `json:"language"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

func (SourceCode) TableName() string {
	return "source_codes"
}
