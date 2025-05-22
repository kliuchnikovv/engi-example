package entity

import "time"

type Note struct {
	ID        int64     `json:"id"                gorm:"primaryKey;autoIncrement" description:"Unique note identifier"           example:"1"`
	Title     string    `json:"title"             gorm:"column:title"             description:"Note title"                       example:"Meeting notes"`
	Content   string    `json:"content,omitempty" gorm:"column:content"           description:"Note content in Markdown"         example:"# Heading"`
	CreatedAt time.Time `json:"created_at"        gorm:"column:created_at"        description:"Note creation timestamp (UTC)"    example:"2024-06-01T12:00:00Z"`
	UpdatedAt time.Time `json:"updated_at"        gorm:"column:updated_at"        description:"Note last update timestamp (UTC)" example:"2024-06-01T12:00:00Z"`
}
