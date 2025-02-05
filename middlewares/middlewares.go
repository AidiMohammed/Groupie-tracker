package middlewares

import (
	"net/http"
	"fmt"
	"groupie-tracker/controllers/system"
	
)
// Middleware to handle unexpected server errors
func withErrorHandling(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// error in server
				controller.ErrorPage(w, http.StatusInternalServerError, "Server encountered an unexpected error")
				fmt.Println("Recovered from panic:", rec)
			}
		}()
		next(w, r)
	}
}
