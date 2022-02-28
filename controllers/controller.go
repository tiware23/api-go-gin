package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tiware23/api-go-gin/models"
)

func ExibirTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E ai" + nome + ", td blz?",
	})
}
