package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/ukasyah99/user-api/model"
)

type UserCreateBody struct {
	DisplayName string `json:"display_name" binding:"required"`
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserCreate godoc
// @Description Create a new user
// @Accept json
// @Produce json
// @Param Authorization header string true "As bearer token, e.g. 'Bearer <access-token>'"
// @Param display_name body string true "Display name"
// @Param name body string true "Name"
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} UserResponse
// @Router /users [post]
func UserCreate(db *sqlx.DB, c *gin.Context) {
	var body UserCreateBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO users (display_name, name, email, password) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, body.DisplayName, body.Name, body.Email, body.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var id int64
	id, err = result.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	user := model.User{}

	err = db.Get(&user, "SELECT * FROM users WHERE id = ? LIMIT 1", id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}
