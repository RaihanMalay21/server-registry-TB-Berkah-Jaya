package main

import (
	"github.com/gorilla/mux"
	// "github.com/gorilla/handlers"
	"net/http"
	"log"
	"fmt"

	"github.com/RaihanMalay21/server-registry-TB-Berkah-Jaya/controller"
	"github.com/RaihanMalay21/server-registry-TB-Berkah-Jaya/controller/template"
	config "github.com/RaihanMalay21/config-tb-berkah-jaya"
)

func main() {
	r := mux.NewRouter()

	config.DB_Connection()
	r.Use(corsMiddlewares)
	api := r.PathPrefix("/berkahjaya").Subrouter()
	api.HandleFunc("/login", controller.Login).Methods("POST")
	api.HandleFunc("/signup", controller.SignUp).Methods("POST") 
	api.HandleFunc("/logout", controller.LogOut).Methods("GET")
	// r.HandleFunc("/get/hadiah", controller.Hadiah).Methods("GET")
	api.HandleFunc("/forgot/password", controller.ForgotPassword).Methods("POST")
	api.HandleFunc("/forgot/password/reset", template.PageResetPassword).Methods("GET")
	api.HandleFunc("/forgot/password/reset", controller.ForgotPasswordChangePassword).Methods("POST")

	// corsHandler := handlers.CORS(
	// 	handlers.AllowedOrigins([]string{"https://fe-tb-berkah-jaya-igcfjdj5fa-uc.a.run.app"}),
	// 	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	// 	handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	// 	handlers.AllowCredentials(),
	// )

	log.Fatal(http.ListenAndServe(":8080", r))
}

func corsMiddlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Println("Origin received:", origin)

		allowedOrigins := "http://localhost:3000"

		if origin == allowedOrigins {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}