package controllers

import (
	// "cronbackend/db"
	// "cronbackend/models"
	// "cronbackend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "net/http"
	// "strings"
	"fmt"
)

type CheckingController struct {
	DB *gorm.DB
}

func NewCheckingController(db *gorm.DB) *CheckingController {
	return &CheckingController{DB: db}
}

func (cc *CheckingController) Ping(c *gin.Context) {

	userId := c.Param("userID")
	pingId := c.Param("pingID")

	fmt.Println("userId: ", userId)
	fmt.Println("pingId: ", pingId)

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
