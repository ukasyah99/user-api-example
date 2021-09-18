package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type AuthLogoutBody struct {
	RefreshToken string `json:"refresh_token" binding:"required" example:"BpLnfgDsc2WD8F2qNfHK5a84jjJkwzDk"`
}

type AuthLogoutResponseData struct {
	Message string `json:"message" example:"Logged out successfully"`
}

type AuthLogoutResponse struct {
	Data AuthLogoutResponseData `json:"data"`
}

// AuthLogout godoc
// @Description Logout by deleting refresh token
// @Accept json
// @Produce json
// @Param Authorization header string true "As bearer token, e.g. 'Bearer <access-token>'"
// @Param refresh_token body string true "This token is obtained when login successfully"
// @Success 200 {object} AuthLogoutResponse
// @Router /auth/logout [post]
func AuthLogout(db *sqlx.DB, c *gin.Context) {
	var body AuthLogoutBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query := "DELETE FROM refresh_tokens WHERE id = ?"
	if _, err := db.Exec(query, body.RefreshToken); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": AuthLogoutResponseData{
			Message: "Logged out successfully",
		},
	})
}
