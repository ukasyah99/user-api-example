package handler

import (
    "github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/ukasyah99/user-api/model"
)

type UserResponse struct {
	Data model.User `json:"data"`
}

// AuthMe godoc
// @Description Get current user
// @Accept json
// @Produce json
// @Param Authorization header string true "As bearer token, e.g. 'Bearer <access-token>'"
// @Success 200 {object} UserResponse
// @Router /auth/me [get]
func AuthMe(db *sqlx.DB, c *gin.Context) {
	user := c.MustGet("user").(model.User)
	c.JSON(200, gin.H{
		"data": user,
	})
}
