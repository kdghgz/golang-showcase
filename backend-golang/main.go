package main

import (
	"santrikoding/backend-api/config"
	"santrikoding/backend-api/database"
	"santrikoding/backend-api/routes"
	"time"

	"github.com/gin-contrib/cors"
)

func main() {
	// Load config .env
	config.LoadEnv()

	// Inisialisasi database
	database.InitDB()

	// Setup router
	r := routes.SetupRouter()

	// ---- MULAI KODE CORS (SISIPKAN DI SINI) ----
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// ---- SELESAI KODE CORS ----

	// Mulai server
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
