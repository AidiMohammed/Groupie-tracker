package routes

import (
	"groupie-tracker/controllers"
	"groupie-tracker/middlewares"
	"net/http"
	//"fmt"
)

func RegisterRoutes(mux *http.ServeMux){
	//mux.HandleFunc("/", middlewares.withErrorHandling(controller.GetArtists))
	controller.GetArtists(w httmp.)
}