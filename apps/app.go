package apps

import (
	"log"
	"os"

	"github.com/autumnleaf-ra/phonebook-api/database"
	"github.com/autumnleaf-ra/phonebook-api/models"
	"github.com/autumnleaf-ra/phonebook-api/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func InitializeApp(app *fiber.App) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := database.Connect(config)

	if err != nil {
		log.Fatal("Could not load database !")
	}

	err = models.MigrateUser(db)

	if err != nil {
		log.Fatal("Could not migrates db")
	}

	repo := repository.Repository{
		DB: db,
	}

	app.Use(cors.New(cors.Config{AllowCredentials: true}))
	repo.SetupRoute(app)
	app.Listen(":" + port)
}
