package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/ukasyah99/user-api/model"
)

// UserFindOne godoc
// @Description Get specific user
// @Accept json
// @Produce json
// @Param Authorization header string true "As bearer token, e.g. 'Bearer <access-token>'"
// @Param id path int true "Id"
// @Success 200 {object} UserResponse
// @Router /users/:id [get]
func UserFindOne(db *sqlx.DB, c *gin.Context) {
	user := model.User{}

	err := db.Get(&user, "SELECT * FROM users WHERE id = ? LIMIT 1", c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}
