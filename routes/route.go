package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiware23/api-go-gin/controllers"
)

func HandlerRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibirTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/alunos/rg/:rg", controllers.BuscaAlunoPorRG)
	r.Run()
}
