package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func CreateDeck(c *gin.Context) {
	uuid, err := uuid.NewV4()

	if err != nil {
		log.Printf("Something went wrong: %s", err)
		return
	}

}

func ShowDecks(c *gin.Context) {
	c.JSON(200, gin.H{
		"value": 200,
	})
}
