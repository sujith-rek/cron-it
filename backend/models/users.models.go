package models

type User struct {
	ID            string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Email         string          `gorm:"uniqueIndex;not null" json:"email"`
	Password      string          `gorm:"not null" json:"password_hash"`
	Name          string          `json:"name"`
	Limit         int             `json:"limit" gorm:"default:3"`
	ScheduledJobs []ScheduledJobs `gorm:"foreignKey:UserID" json:"scheduled_jobs"`
	CheckJobs     []CheckJobs     `gorm:"foreignKey:UserID" json:"check_jobs"`
}