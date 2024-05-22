package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"pencils-up/github.com/internal/handlers"

	"github.com/joho/godotenv"
)

func main() {
	//Load Environment Varaiables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	http.HandleFunc("GET /", handlers.HomePageHandler)
	http.HandleFunc("GET /login", handlers.LoginPageHandler)
	http.HandleFunc("GET /sign-up", handlers.SignUpPageHandler)
	http.HandleFunc("GET /admin-dashboard", handlers.AdminDashboardPageHandler)

	http.ListenAndServe(port, nil)
	fmt.Println(port)

}
