package models

type User struct {
	ID            string        `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	Email         string        `gorm:"uniqueIndex;not null" json:"email"`
	Password      string        `gorm:"not null" json:"password_hash"`
	Name          string        `json:"name" gorm:"not null"`
	CheckLimit    int           `json:"check_limit" gorm:"default:3"`
	ScheduleLimit int           `json:"schedule_limit" gorm:"default:3"`
	ScheduledJobs []ScheduleJob `gorm:"foreignKey:UserID" json:"scheduled_jobs"`
	CheckJobs     []CheckJob    `gorm:"foreignKey:UserID" json:"check_jobs"`
}

type UserSignUp struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	CheckLimit    int    `json:"check_limit"`
	ScheduleLimit int    `json:"schedule_limit"`
}
