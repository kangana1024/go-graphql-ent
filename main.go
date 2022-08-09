package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kangana1024/go-graphql-ent/database"
	"github.com/kangana1024/go-graphql-ent/ent"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if os.Getenv("CHECK_ENV") != "1" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error : %+v!", err)
		}
	}
	e := echo.New()

	var entOps []ent.Option
	entOps = append(entOps, ent.Debug())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		})
	})
	err := database.NewPGConnectionDB()
	if err != nil {
		log.Fatalf("Database Connection Error : %v\n", err)
	}

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_URL")))
}
