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
	r.Run()
}
