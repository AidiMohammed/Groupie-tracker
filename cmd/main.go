package main

import (
	"fmt"
	"groupie-tracker/config"
	"groupie-tracker/routes"
	//"net/http"

	//controllers "groupie-tracker/back-end/Controllers"
)

func main() {
	config,err := config.LoadConfig();
	if err != nil {
		fmt.Printf("imposible de chager les varaibles d'enviranment : %v",err)
	}
 
	mux := http.NewServerMux()
	routes.RegisterRoutes(mux)

	fmt.Println("starting server on port :"+config.APP_PORT+" ... : "+config.BASE_URL+":"+config.APP_PORT)
	if err := http.ListenAndServe("localhost:8082", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
