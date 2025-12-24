package timelog

import "time"

type Task struct {
	ID          uint    `json:"id"          gorm:"primaryKey"`
	Title       string  `json:"title"`
	Description *string `json:"description"`

	GroupID uint  `json:"group_id"`
	Group   Group `json:"group"    gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Deleted bool `json:"deleted"`

	CreatedTime time.Time  `json:"created_time"`
	DeletedTime *time.Time `json:"deleted_time"`
	DueTime     *time.Time `json:"due_time"`
}

type Group struct {
	ID    uint   `json:"id"    gorm:"primaryKey"`
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}
