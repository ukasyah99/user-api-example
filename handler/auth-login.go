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

type AuthLoginBody struct {
	Name string `json:"name" binding:"required" example:"ukasyah"`
	Password string `json:"password" binding:"required" example:"rahasia"`
}

type AuthLoginResponseData struct {
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2MzE5NDYzMDEsImlzcyI6ImdvLXVzZXItYXBpIn0.bR7ooUYtwG7xocBZgNJFu8l4O5rxdawKCYbazXpDZX4"`
	RefreshToken string `json:"refresh_token" example:"BpLnfgDsc2WD8F2qNfHK5a84jjJkwzDk"`
}

type AuthLoginResponse struct {
	Data AuthLoginResponseData `json:"data"`
}

// AuthLogin godoc
// @Description Login using name and password
// @Accept json
// @Produce json
// @Param name body string true "Username, e.g. ukasyah99"
// @Param password body string true "Password, e.g. rahasia"
// @Success 200 {object} AuthLoginResponse
// @Router /auth/login [post]
func AuthLogin(db *sqlx.DB, c *gin.Context) {
	var body AuthLoginBody
	var query string

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := model.User{}

	query = "SELECT * FROM users WHERE name = ? LIMIT 1"
	if err := db.Get(&user, query, body.Name); err != nil {
		c.JSON(401, gin.H{"error": "Invalid name or password"})
		return
	}

	if user.Password != body.Password {
		c.JSON(401, gin.H{"error": "Invalid name or password"})
		return
	}

	refreshToken := lib.GenerateRandomString(32)
	validUntil := time.Now().AddDate(0, 0, 7).Format("2006-01-02 15:04:05")

	accessTokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JwtClaims{
		user.Id,
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

	query = "INSERT INTO refresh_tokens (id, user_id, valid_until) VALUES (?, ?, ?)"
	if _, err := db.Exec(query, refreshToken, user.Id, validUntil); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": AuthLoginResponseData{
			AccessToken: accessToken,
			RefreshToken: refreshToken,
		},
	})
}
