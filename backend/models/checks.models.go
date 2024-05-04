package models

import (
	"encoding/json"
	"time"
)

type CheckJobs struct {
	ID         string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID     string          `gorm:"type:uuid;not null" json:"user_id"`
	JobName    string          `json:"job_name" gorm:"not null"`
	Url        string          `json:"url"`
	Headers    json.RawMessage `json:"headers"`
	CronString string          `json:"cron_string" gorm:"not null"`
	CreatedAt  time.Time       `json:"created_at" gorm:"default:now()"`
	UpdatedAt  time.Time       `json:"updated_at" gorm:"default:now()"`
}
