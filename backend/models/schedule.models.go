package models

import (
	"encoding/json"
	"time"
)

type ScheduleJob struct {
	ID               string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	UserID           string          `gorm:"type:uuid;not null" json:"user_id"`
	ClusterID        string          `gorm:"type:uuid;not null" json:"cluster_id"`
	ExecString       string          `json:"exec_string" gorm:"not null"`
	AdditionalParams json.RawMessage `json:"additional_params"`
	URL              string          `json:"url" gorm:"not null"`
	CreatedAt        time.Time       `json:"created_at" gorm:"default:now()"`
	UpdatedAt        time.Time       `json:"updated_at" gorm:"default:now()"`
	Name             string          `json:"name" gorm:"not null"`
	
	Cluster          ScheduleCluster `gorm:"foreignKey:ClusterID"`
}

type ScheduleCluster struct {
	ID              string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	Name            string    `json:"name" gorm:"not null"`
	ExecutionString string    `json:"execution_string" gorm:"not null"`
	Size            int       `json:"size" gorm:"default:0"`
	CreatedAt       time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"default:now()"`

	Jobs []ScheduleJob `gorm:"foreignKey:ClusterID" json:"jobs"`
}

type JobInput struct {
	Name             string          `json:"name" binding:"required"`
	ExecString       string          `json:"exec_string" binding:"required"`
	AdditionalParams json.RawMessage `json:"additional_params"`
	URL              string          `json:"url" binding:"required"`
}
