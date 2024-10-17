package main

import (
	"log"

	"github.com/dzakimaulana/SiJaki-Backend/internal/config"
	"github.com/dzakimaulana/SiJaki-Backend/internal/database"
	"github.com/dzakimaulana/SiJaki-Backend/internal/handlers"
	"github.com/dzakimaulana/SiJaki-Backend/internal/services"
)

func main() {
	cfg := config.LoadConfig()

	dbConn, err := database.ConnectDB(cfg.DBUser, cfg.DBName, cfg.DBPassword)
	if err != nil {
		log.Fatalf("😞 Failed to connect to the database: %v", err)
	}
	log.Println("✅ Successfully connected to the database!")
	defer closeDB(dbConn)

	userSvc := services.NewGeneralSvc(dbConn)
	workerSvc := services.NewWorkerSvc(dbConn)

	userHandler := handlers.NewUserHandler()
	workerHandler := handlers.NewWorkerHandler()

}
