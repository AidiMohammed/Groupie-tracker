package controller

import (
	"groupie-tracker/tools"
	"html/template"
	"net/http"
)

type ErrorData struct {
	CodeStatus	int
	Message		string
}

func ErrorPage(w http.ResponseWriter, statuCode int,message string) {
	if tools.FileExists("../views/error.html") {
		template,err := template.ParseFiles("../views/error.html")

		if err != nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}

		w.WriteHeader(statuCode)
		data := ErrorData{
			CodeStatus: statuCode,
			Message: message,
		}
		template.ExecuteTemplate(w,"error.html", data)
	}
}