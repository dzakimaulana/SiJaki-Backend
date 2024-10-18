package main

import (
	"log"

	"github.com/dzakimaulana/SiJaki-Backend/internal/config"
	"github.com/dzakimaulana/SiJaki-Backend/internal/database"
	"github.com/dzakimaulana/SiJaki-Backend/internal/handlers"
	"github.com/dzakimaulana/SiJaki-Backend/internal/middlewares"
	"github.com/dzakimaulana/SiJaki-Backend/internal/routes"
	"github.com/dzakimaulana/SiJaki-Backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfgDB := config.LoadDBConfig()
	cfgApp := config.LoadAppConfig()

	dbConn, err := database.ConnectDB(cfgDB.DBUser, cfgDB.DBName, cfgDB.DBPassword)
	if err != nil {
		log.Fatalf("😞 Failed to connect to the database: %v", err)
	}
	log.Println("✅ Successfully connected to the database!")

	app := fiber.New()

	api := app.Group("/api")
	apiUser := api.Group("/users")
	apiWorker := api.Group("/workers", middlewares.OnlyAdmin)

	userSvc := services.NewUserSvc(dbConn)
	workerSvc := services.NewWorkerSvc(dbConn)

	userHandler := handlers.NewUserHandler(userSvc)
	workerHandler := handlers.NewWorkerHandler(workerSvc)

	routes.UserRoute(userHandler, apiUser)
	routes.WorkerRoute(workerHandler, apiWorker)

	app.Listen(cfgApp.Port)
}
