package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/ukasyah99/user-api/model"
)

// UserDelete godoc
// @Description Delete existing user
// @Accept json
// @Produce json
// @Param Authorization header string true "As bearer token, e.g. 'Bearer <access-token>'"
// @Param id path int true "Id"
// @Success 200 {object} UserResponse
// @Router /users/:id [delete]
func UserDelete(db *sqlx.DB, c *gin.Context) {
	var query string

	user := model.User{}

	query = "SELECT * FROM users WHERE id = ? LIMIT 1"
	if err := db.Get(&user, query, c.Param("id")); err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	query = "DELETE FROM users WHERE id = ?"
	if _, err := db.Exec(query, c.Param("id")); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}
