package main

import (
	"fmt"
	"github.com/spoonboy-io/todo-deployment-app/internal/postgres"
	//"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
	"github.com/spoonboy-io/todo-deployment-app/internal/routes"
	//"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/spoonboy-io/koan"
)

const (
	CONFIG_FILE = "config.env"
	PG_USER="postgres"
	PG_PASSWORD="Password123?"
	PG_DATABASE="todos"
	APP_PORT_DEFAULT = "8090"
)

var logger *koan.Logger

func init(){
	// read in the config file
	err := godotenv.Load(CONFIG_FILE)
	if err != nil {
		logger.FatalError("Failed to read config file", err)
	}
}

func main() {
	// connect to database
	dbHost := os.Getenv("PG_HOST")
	if dbHost == "" {
		logger.FatalError("No database host found", nil)
	}

	dbUser := os.Getenv("PG_USER")
	if dbUser == ""{
		dbUser = PG_USER
	}
	dbPassword := os.Getenv("PG_PASSWORD")
	if dbPassword == "" {
		dbPassword = PG_PASSWORD
	}
	dbDatabase := os.Getenv("PG_DATABASE")
	if dbDatabase == ""{
		dbDatabase = PG_DATABASE
	}

	db, err := postgres.Connect(dbHost, dbUser, dbPassword, dbDatabase)
	if err != nil {
		logger.FatalError("Failed to make database connection", err)
	}

	engine :=html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// routes
	app.Get("/", func(c *fiber.Ctx) error {
		return routes.IndexHandler(c, db, logger)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return routes.PostHandler(c, db, logger)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return routes.PutHandler(c, db, logger)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return routes.DeleteHandler(c, db, logger)
	})

	// start server
	appPort := os.Getenv("APP_SERVER_PORT")
	if appPort == "" {
		appPort = APP_PORT_DEFAULT
	}
	if err := app.Listen(fmt.Sprintf(":%v", appPort)); err != nil {
		logger.FatalError("Failed to start server", err)
	}
}