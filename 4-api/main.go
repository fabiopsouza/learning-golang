package main

import (
	"fmt"
	"github.com/fabipereira/learning-golang/4-api/models"
	"github.com/fabipereira/learning-golang/4-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	routes.HandleRequest()
	fmt.Println("Server started")
}
