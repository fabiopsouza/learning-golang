package routes

import (
	"github.com/fabipereira/learning-golang/4-api/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	//r := mux.NewRouter()

	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
