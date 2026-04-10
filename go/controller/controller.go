package controller

import (
	"fmt"
	"tcg_deck/go/model"
	"tcg_deck/go/usecase"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	u usecase.Usecase
}

func NewController(u usecase.Usecase) Controller {
	return Controller{u: u}
}

func (c *Controller) GetDeck(g gin.Context) {
	idParams := g.Param("id")
	var id int32
	_, err := fmt.Sscanf(idParams, "id", &id)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := c.u.GetDeck(id)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, result)
}

func (c *Controller) CreateDeck(g gin.Context) {
	var response model.Response
	if err := g.ShouldBindJSON(&response); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := c.u.CreateDeck(response)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, result)
}
