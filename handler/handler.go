package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type HandlerFunction func(db *sqlx.DB, c *gin.Context)
