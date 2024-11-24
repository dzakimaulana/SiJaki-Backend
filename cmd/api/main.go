package main

import (
	"fmt"
	"log"

	"github.com/dzakimaulana/SiJaki-Backend/internal/config"
	"github.com/dzakimaulana/SiJaki-Backend/internal/database"
	"github.com/dzakimaulana/SiJaki-Backend/internal/handlers"
	"github.com/dzakimaulana/SiJaki-Backend/internal/middlewares"
	"github.com/dzakimaulana/SiJaki-Backend/internal/routes"
	"github.com/dzakimaulana/SiJaki-Backend/internal/services"
	"github.com/dzakimaulana/SiJaki-Backend/internal/transport"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
)

func main() {
	// Config
	cfgDB := config.LoadDBConfig()
	cfgApp := config.LoadAppConfig()
	cfgMqtt := config.LoadMqttConfig()

	// MQTT
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", cfgMqtt.Broker, cfgMqtt.Port)) // it musth be changed if we use name server to redirect the connection
	opts.SetUsername(cfgMqtt.Username)
	opts.SetPassword(cfgMqtt.Password)
	opts.SetClientID("BE")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// Database PostgreSQL
	dbConn, err := database.ConnectDB(cfgDB.DBUser, cfgDB.DBName, cfgDB.DBPassword)
	if err != nil {
		log.Fatalf("😞 Failed to connect to the database: %v", err)
	}
	log.Println("✅ Successfully connected to the database!")

	// Backend
	app := fiber.New()

	app.Use(logger.New())

	// WebSocket route
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		transport.Subscribe(client, c)
		transport.WebSocketHandler(c)
	}))

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
