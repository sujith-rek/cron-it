package middleware

import (
	"cronbackend/db"
	"cronbackend/models"
	"cronbackend/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// TokenResponse represents the structure of the access and refresh token response.
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func UnwrapUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accessToken, refreshToken string

		// Retrieve tokens from cookies or Authorization header
		accessTokenCookie, err := c.Cookie("access_token")
		refreshTokenCookie, _ := c.Cookie("refresh_token")

		authorizationHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) == 2 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = accessTokenCookie
		}

		if accessToken == "" && refreshTokenCookie == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, err := db.LoadConfig(".")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Server error"})
			return
		}

		if accessToken == "" && refreshTokenCookie != "" {
			refreshToken = refreshTokenCookie
			newAccessToken, newRefreshToken, userResponse, tokenErr := utils.GenerateNewTokens(refreshToken, config)

			if tokenErr != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid or expired refresh token"})
				return
			}

			if userResponse == nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid or expired refresh token"})
				return
			}

			userInString, _ := json.Marshal(userResponse)

			accessToken = newAccessToken

			// Set the new access token and refresh token as cookies
			c.SetCookie("access_token", newAccessToken, config.AccessTokenMaxAge*60+19800, "/", "", false, true)
			c.SetCookie("refresh_token", newRefreshToken, config.RefreshTokenMaxAge*60+19800, "/", "", false, true)
			c.SetCookie("user", string(userInString), config.AccessTokenMaxAge*60+19800, "/", "", false, false)

			// Update the access token for further processing
			accessToken = newAccessToken
		}

		sub, validateErr := utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
		if validateErr != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid or expired access token"})
			return
		}

		// Token is valid, retrieve the user
		subMap := sub.(map[string]interface{})
		userID, ok := subMap["id"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid token data"})
			return
		}

		var user models.User
		result := db.DB.First(&user, "id = ?", fmt.Sprint(userID))
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "The user belonging to this token no longer exists"})
			return
		}

		// Set user in context for further use in the request pipeline
		c.Set("user", user)
		c.Next()
	}
}
