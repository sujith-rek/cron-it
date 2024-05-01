package models

import (
	"encoding/json"
	"time"
)

// User table
type User struct {
	ID            string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Email         string          `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash  string          `gorm:"not null" json:"password_hash"`
	Name          string          `json:"name"`
	Limit         int             `json:"limit"`
	ScheduledJobs []ScheduledJobs `gorm:"foreignKey:UserID" json:"scheduled_jobs"`
	CheckJobs     []CheckJobs     `gorm:"foreignKey:UserID" json:"check_jobs"`
}

// Job table that are scheduled to run
type ScheduledJobs struct {
	ID              string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID          string          `gorm:"type:uuid;not null" json:"user_id"`
	JobName         string          `json:"job_name"`
	CronString      string          `json:"cron_string"`
	Url             string          `json:"url"`
	AdditonalParams json.RawMessage `json:"additonal_params"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

// Jobs to check for
type CheckJobs struct {
	ID         string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID     string    `gorm:"type:uuid;not null" json:"user_id"`
	JobName    string    `json:"job_name"`
	Url        string    `json:"url"`
	CronString string    `json:"cron_string"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
