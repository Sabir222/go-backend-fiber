package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sabir222/go-backend-fiber/internal/server"
	"os"
	"strconv"
)

func main() {

	server := server.New()

	server.SetupRoutes()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	fmt.Println(port)
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
