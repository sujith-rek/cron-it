package models

import (
	"encoding/json"
	"time"
)

type ScheduledJobs struct {
	ID              string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID          string          `gorm:"type:uuid;not null" json:"user_id"`
	JobName         string          `json:"job_name" gorm:"not null"`
	CronString      string          `json:"cron_string" gorm:"not null"`
	Url             string          `json:"url" gorm:"not null"`
	AdditonalParams json.RawMessage `json:"additonal_params"`
	CreatedAt       time.Time       `json:"created_at" gorm:"default:now()"`
	UpdatedAt       time.Time       `json:"updated_at" gorm:"default:now()"`
}
