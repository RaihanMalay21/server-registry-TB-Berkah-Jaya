package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"

	"github.com/RaihanMalay21/server-registry-TB-Berkah-Jaya/controller"
	"github.com/RaihanMalay21/server-registry-TB-Berkah-Jaya/controller/template"
	config "github.com/RaihanMalay21/config-TB_Berkah_Jaya"
)

func main() {
	r := mux.NewRouter()

	config.DB_Connection()
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/signup", controller.SignUp).Methods("POST") 
	r.HandleFunc("/logout", controller.LogOut).Methods("GET")
	// r.HandleFunc("/get/hadiah", controller.Hadiah).Methods("GET")
	r.HandleFunc("/forgot/password", controller.ForgotPassword).Methods("POST")
	r.HandleFunc("/forgot/password/reset", template.PageResetPassword).Methods("GET")
	r.HandleFunc("/forgot/password/reset", controller.ForgotPasswordChangePassword).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}