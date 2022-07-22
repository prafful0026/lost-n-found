package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (cnt *Controllers) HandlePing(c *gin.Context) {
	uID := (c.MustGet("id").(string))
	log.Println(uID)
	c.JSON(200, gin.H{
		"message": uID,
	})
}
