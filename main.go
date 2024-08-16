package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"
	"log"
	"os"

	"github.com/RaihanMalay21/server-registry-TB-Berkah-Jaya/controller"
	"github.com/RaihanMalay21/server-registry-TB-Berkah-Jaya/controller/template"
	config "github.com/RaihanMalay21/config-tb-berkah-jaya"
)

func main() {
	r := mux.NewRouter()

	config.DB_Connection()
	api := r.PathPrefix("/berkahjaya").Subrouter()
	api.HandleFunc("/login", controller.Login).Methods("POST")
	api.HandleFunc("/signup", controller.SignUp).Methods("POST") 
	api.HandleFunc("/logout", controller.LogOut).Methods("GET")
	// r.HandleFunc("/get/hadiah", controller.Hadiah).Methods("GET")
	api.HandleFunc("/forgot/password", controller.ForgotPassword).Methods("POST")
	api.HandleFunc("/forgot/password/reset", template.PageResetPassword).Methods("GET")
	api.HandleFunc("/forgot/password/reset", controller.ForgotPasswordChangePassword).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", 
		handlers.CROS(
			handlers.AllowedOrigins([]string{"http://localhost:3000"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedCredentials(),
		)(r)))
}