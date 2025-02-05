package main

import (
	"fmt"
	"groupie-tracker/config"
	"groupie-tracker/routes"
	"net/http")

func main() {
	config,err := config.LoadConfig();
	if err != nil {
		fmt.Printf("imposible de chager les varaibles d'enviranment : %v",err)
	}
 
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	fmt.Println("starting server on port :"+config.AppPort+" ... : "+config.BaseUrl+":"+config.AppPort)
	if err := http.ListenAndServe("localhost:"+config.AppPort, mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
