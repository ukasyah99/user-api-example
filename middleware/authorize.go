package middleware

import (
    "github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/ukasyah99/user-api/model"
    "os"
    "strings"
)

func Authorize(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeaders := c.Request.Header["Authorization"]
		if len(authHeaders) == 0 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Access token must be provided"})
			return
		}
	
		accessToken := strings.Replace(authHeaders[0], "Bearer ", "", 1)

		token, err := jwt.ParseWithClaims(accessToken, &model.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return	
		}

		if claims, ok := token.Claims.(*model.JwtClaims); ok && token.Valid {
			user := model.User{}

			err := db.Get(&user, "SELECT * FROM users WHERE id = ? LIMIT 1", claims.UserId)
			if err != nil {
				c.JSON(500, gin.H{"error": "User not found"})
				return
			}

			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatusJSON(401, gin.H{
				"error": err.Error(),
			})
		}
	}
}
