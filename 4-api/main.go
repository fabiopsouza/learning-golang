package main

import (
	"fmt"
	"github.com/fabiopsouza/learning-golang/4-api/models"
	"github.com/fabiopsouza/learning-golang/4-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	routes.HandleRequest()
	fmt.Println("Server started")
}
