package main

import (
	"github.com/tiware23/api-go-gin/database"
	"github.com/tiware23/api-go-gin/models"
	"github.com/tiware23/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Thiago", CPF: "00000000", RG: "470000000"},
		{Nome: "Fernanda", CPF: "11000000", RG: "480000000"},
	}
	routes.HandlerRequest()
}
