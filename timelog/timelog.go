package timelog

import "time"

type Task struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Group       string  `json:"group"`

	Deleted bool `json:"deleted"`

	CreatedTime time.Time  `json:"created_time"`
	DeletedTime *time.Time `json:"deleted_time"`
	DueTime     *time.Time `json:"due_time"`
}
