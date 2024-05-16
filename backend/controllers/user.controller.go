package controllers

import (
	"cronbackend/db"
	"cronbackend/models"
	"cronbackend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DataBase *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DataBase: db}
}

func (uc *UserController) Register(c *gin.Context) {

	var userSignUp models.UserSignUp

	if err := c.ShouldBindJSON(&userSignUp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if strings.TrimSpace(userSignUp.Email) == "" || strings.TrimSpace(userSignUp.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	hashedPassword, err := utils.HashPassword(userSignUp.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password"})
		return
	}

	user := models.User{
		Email:    userSignUp.Email,
		Password: hashedPassword,
		Name:     userSignUp.Name,
		Limit:    3,
	}

	res := uc.DataBase.Create(&user)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error(), "message": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})
}

func (uc *UserController) Login(c *gin.Context) {
	var userLogin models.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if strings.TrimSpace(userLogin.Email) == "" || strings.TrimSpace(userLogin.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	var user models.User

	uc.DataBase.Where("email = ?", userLogin.Email).First(&user)

	if user.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if err := utils.VerifyPassword(user.Password, userLogin.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	userResponse := models.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Limit: user.Limit,
	}

	config, _ := db.LoadConfig(".")

	a_token, err := utils.CreateToken(config.AccessTokenExpiresIn, userResponse, config.AccessTokenPrivateKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	r_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, userResponse, config.RefreshTokenPrivateKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	userInString, _ := json.Marshal(userResponse)

	c.SetCookie("access_token", a_token, config.AccessTokenMaxAge*60 + 19800, "/", "", false, true)
	c.SetCookie("refresh_token", r_token, config.RefreshTokenMaxAge*60 + 19800, "/", "", false, true)
	c.SetCookie("user",string(userInString) , config.AccessTokenMaxAge*60 + 19800, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{"message": "user logged in", "user": userResponse})

}

func (uc *UserController) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no refresh token found"})
		return
	}

	config, _ := db.LoadConfig(".")

	user, err := utils.ValidateToken(refreshToken, config.RefreshTokenPublicKey)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	userResponse := user.(models.UserResponse)

	a_token, err := utils.CreateToken(config.AccessTokenExpiresIn, userResponse, config.AccessTokenPrivateKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	r_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, userResponse, config.RefreshTokenPrivateKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.SetCookie("access_token", a_token, config.AccessTokenMaxAge*60 + 19800, "/", "", false, true)
	c.SetCookie("refresh_token", r_token, config.RefreshTokenMaxAge*60 + 19800, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "token refreshed"})
}

func (uc *UserController) Logout(c *gin.Context) {

	// get the user from the context
	user, _ := c.Get("user")
	fmt.Println(user)

	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "user logged out"})
}
