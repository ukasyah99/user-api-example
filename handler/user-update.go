package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/ukasyah99/user-api/model"
)

type UserUpdateBody struct {
	DisplayName string `json:"display_name" binding:""`
	Name string `json:"name" binding:""`
	Email string `json:"email" binding:""`
	Password string `json:"password" binding:""`
}

// UserUpdate godoc
// @Description Update user
// @Accept json
// @Produce json
// @Param Authorization header string true "As bearer token, e.g. 'Bearer <access-token>'"
// @Param id path int true "Id"
// @Param display_name body string false "Display name"
// @Param name body string false "Name"
// @Param email body string false "Email"
// @Param password body string false "Password"
// @Success 200 {object} UserResponse
// @Router /users/:id [put]
func UserUpdate(db *sqlx.DB, c *gin.Context) {
	var body UserUpdateBody
	var query string

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := model.User{}

	query = "SELECT * FROM users WHERE id = ? LIMIT 1"
	if err := db.Get(&user, query, c.Param("id")); err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "User not found"})
		return
	}

	if body.DisplayName == "" { body.DisplayName = user.DisplayName }
	if body.Name == "" { body.Name = user.Name }
	if body.Email == "" { body.Email = user.Email }
	if body.Password == "" { body.Password = user.Password }

	query = "UPDATE users SET display_name = ?, name = ?, email = ?, password = ? WHERE id = ?"
	_, err := db.Exec(query, body.DisplayName, body.Name, body.Email, body.Password, c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	query = "SELECT * FROM users WHERE id = ? LIMIT 1"
	if err := db.Get(&user, query, c.Param("id")); err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}
