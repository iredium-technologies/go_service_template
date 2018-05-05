package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iredium-technologies/go_service_template/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()

	routes.Routes(e)

	listen := fmt.Sprintf(":%s", os.Getenv("PORT"))
	e.Logger.Fatal(e.Start(listen))
}
