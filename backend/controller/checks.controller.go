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

type CheckingController struct {
	DB *gorm.DB
}

func NewCheckingController(db *gorm.DB) *CheckingController {
	return &CheckingController{DB: db}
}