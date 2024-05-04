package controller

import (
	// "cronbackend/db"
	// "cronbackend/models"
	// "cronbackend/utils"
	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "net/http"
	// "strings"
)

type ScheduleController struct {
	DB *gorm.DB
}

func NewScheduleController(db *gorm.DB) *ScheduleController {
	return &ScheduleController{DB: db}
}


