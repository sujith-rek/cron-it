package models

import (
	"encoding/json"
	"time"
)

type CheckJob struct {
	ID         string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique" json:"id"`
	UserID     string          `gorm:"type:uuid;not null" json:"user_id"`
	UserMail   string          `json:"user_mail"`
	ClusterID  string          `gorm:"type:uuid;not null" json:"cluster_id"`

	Name       string          `json:"name" gorm:"not null"`
	Body       json.RawMessage `json:"body"`
	PingString string          `json:"ping_string" gorm:"not null"`
	ExecString string          `json:"exec_string" gorm:"not null;default:ping"`

	CreatedAt  time.Time       `json:"created_at" gorm:"default:now()"`
	UpdatedAt  time.Time       `json:"updated_at" gorm:"default:now()"`

	Cluster    CheckCluster    `gorm:"foreignKey:ClusterID"`
}

type CheckCluster struct {
	ID              string     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique" json:"id"`
	Name            string     `json:"name" gorm:"not null"`
	ExecutionString string     `json:"execution_string" gorm:"not null"`
	Size            int        `json:"size" gorm:"not null"`
	CreatedAt       time.Time  `json:"created_at" gorm:"default:now()"`
	LastCheck       time.Time  `json:"last_check" gorm:"default:now()"`
	Jobs            []CheckJob `gorm:"foreignKey:ClusterID" json:"jobs"`
}

