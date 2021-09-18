package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/ukasyah99/user-api/model"
)

type UsersResponse struct {
	Data []model.User `json:"data"`
}

// UserFindAll godoc
// @Description Get all users
// @Accept json
// @Produce json
// @Param Authorization header string true "As bearer token, e.g. 'Bearer <access-token>'"
// @Success 200 {object} UsersResponse
// @Router /users [get]
func UserFindAll(db *sqlx.DB, c *gin.Context) {
	users := []model.User{}

	err := db.Select(&users, "SELECT * FROM users")
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": users,
	})
}
