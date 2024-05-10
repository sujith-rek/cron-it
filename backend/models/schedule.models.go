package models

import (
	"encoding/json"
	"time"
)

// type Job struct {
// 	ID               string
// 	Name             string
// 	UserID           string
// 	ClusterID        string
// 	ExecString       string
// 	AdditionalParams json.RawMessage
// 	URL              string
// 	LastExecuted     string
// 	NextExecution    string
// }

// type Cluster struct {
// 	ID              string
// 	Name            string
// 	ExecutionString string
// 	Jobs            []Job
// 	Size            int
// }

// type ScheduledJobs struct {
// 	ID              string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
// 	UserID          string          `gorm:"type:uuid;not null" json:"user_id"`
// 	JobName         string          `json:"job_name" gorm:"not null"`
// 	CronString      string          `json:"cron_string" gorm:"not null"`
// 	Url             string          `json:"url" gorm:"not null"`
// 	AdditonalParams json.RawMessage `json:"additonal_params"`
// 	CreatedAt       time.Time       `json:"created_at" gorm:"default:now()"`
// 	UpdatedAt       time.Time       `json:"updated_at" gorm:"default:now()"`
// }

type Job struct {
	ID               string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID           string          `gorm:"type:uuid;not null" json:"user_id"`
	ClusterID        string          `gorm:"type:uuid;not null" json:"cluster_id"`
	ExecString       string          `json:"exec_string" gorm:"not null"`
	AdditionalParams json.RawMessage `json:"additional_params"`
	URL              string          `json:"url" gorm:"not null"`
	CreatedAt        time.Time       `json:"created_at" gorm:"default:now()"`
	UpdatedAt        time.Time       `json:"updated_at" gorm:"default:now()"`
	Name             string          `json:"name" gorm:"not null"`

	Cluster Cluster `gorm:"foreignKey:ClusterID"`
}

type Cluster struct {
	ID              string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name            string    `json:"name" gorm:"not null"`
	ExecutionString string    `json:"execution_string" gorm:"not null"`
	Size            int       `json:"size"`
	CreatedAt       time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"default:now()"`
	
	Jobs            []Job     `gorm:"foreignKey:ClusterID" json:"jobs"`

}


type JobInput struct {
	Name             string          `json:"name" binding:"required"`
	ExecString       string          `json:"exec_string" binding:"required"`
	AdditionalParams json.RawMessage `json:"additional_params"`
	URL              string          `json:"url" binding:"required"`
	UserID           string          `json:"user_id" binding:"required"`
}