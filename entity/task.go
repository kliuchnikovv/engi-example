package entity

import "time"

type Task struct {
	ID          int64     `json:"id"                    gorm:"primaryKey;autoIncrement" description:"Unique task identifier"           example:"1"`
	Title       string    `json:"title"                 gorm:"column:title"             description:"Task title"                       example:"Buy groceries"`
	Description string    `json:"description,omitempty" gorm:"column:description"       description:"Task details"                     example:"Milk, eggs, bread"`
	Completed   bool      `json:"completed"             gorm:"column:completed"         description:"Task completion status"           example:"false"`
	CreatedAt   time.Time `json:"created_at"            gorm:"column:created_at"        description:"Task creation timestamp (UTC)"    example:"2024-06-01T12:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at"            gorm:"column:updated_at"        description:"Task last update timestamp (UTC)" example:"2024-06-01T12:00:00Z"`
}
