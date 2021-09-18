package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/ukasyah99/user-api/model"
	"gitlab.com/ukasyah99/user-api/lib"
	"os"
	"time"
)

type AuthRefreshBody struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// AuthRefresh godoc
// @Description Get new access token
// @Accept json
// @Produce json
// @Param refresh_token body string true "This token is obtained when login successfully"
// @Success 200 {object} AuthLoginResponse
// @Router /auth/refresh [post]
func AuthRefresh(db *sqlx.DB, c *gin.Context) {
	var body AuthRefreshBody
	var query string

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if refresh token valid
	refreshTokenRecord := model.RefreshToken{}
	err := db.Get(&refreshTokenRecord, "SELECT * FROM refresh_tokens WHERE id = ? AND valid_until >= CURRENT_TIMESTAMP LIMIT 1", body.RefreshToken)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Refresh token not found or expired"})
		return
	}

	// Generate new access and refresh token
	refreshToken := lib.GenerateRandomString(32)
	validUntil := time.Now().AddDate(0, 0, 7).Format("2006-01-02 15:04:05")
	accessTokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JwtClaims{
		refreshTokenRecord.UserId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			Issuer:    "go-user-api",
		},
	})
	accessToken, err := accessTokenJwt.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Update existing refresh token
	query = "UPDATE refresh_tokens SET id = ?, valid_until = ? WHERE id = ?"
	if _, err := db.Exec(query, refreshToken, validUntil, refreshTokenRecord.Id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"access_token": accessToken,
		"refresh_token": refreshToken,
	})
}
