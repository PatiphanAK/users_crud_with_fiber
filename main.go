package main

import (
	"fmt"
	"log"
	"os"
	"tatar/book/.gen/mydb/database"
	"tatar/book/handler"
	api "tatar/book/oapi"
	"tatar/book/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)
	dbManager, err := database.NewDBManager(dsn)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer dbManager.Close()

	userRepo := dbManager.UserRepository
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	api.RegisterHandlers(app, userHandler)

	app.Listen(":8080")

}
