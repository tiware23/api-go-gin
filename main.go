package main

import (
	"github.com/tiware23/api-go-gin/database"
	"github.com/tiware23/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandlerRequest()
}
