package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/brianvoe/gofakeit/v4"
	"github.com/joho/godotenv"
	"github.com/kangana1024/go-graphql-ent/ent"
	"github.com/kangana1024/go-graphql-ent/ent/schema"
	"github.com/kangana1024/go-graphql-ent/ent/user"
	"github.com/kangana1024/go-graphql-ent/graph"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
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

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panicf("dial database failed: %v", err)
	}
	defer db.Close()
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(0)

	client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, db)))

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		})
	})

	e.POST("/genuser", func(c echo.Context) error {
		a, err := client.Article.Create().SetTitle(gofakeit.Name()).SetDescription(gofakeit.JobDescriptor()).SetContent(&schema.ArticleContent{
			Time: time.Now().Unix(),
			Blocks: []struct {
				Type string      "json:\"type\""
				Data interface{} "json:\"data\""
			}{
				{
					Type: "embded",
					Data: struct {
						Service string `json:"service"`
						Source  string `json:"source"`
						Embed   string `json:"embed"`
						Width   int    `json:"width"`
						Height  int    `json:"height"`
					}{
						Service: "youtube",
						Source:  "https://www.youtube.com/watch?v=JbqGYgI3XY0",
						Embed:   "https://www.youtube.com/embed/JbqGYgI3XY0",
						Height:  320,
					},
				},
			},
			Version: "123.222",
		}).Save(c.Request().Context())

		u, err := client.User.Create().SetBirthDate(gofakeit.Date()).SetFirstname(gofakeit.FirstName()).SetLastname(gofakeit.LastName()).SetUsername(gofakeit.Username()).AddArticles(a).Save(c.Request().Context())
		if !errors.Is(err, nil) {
			return c.JSON(http.StatusInternalServerError, struct {
				Status string `json:"status"`
				Error  string `json:"error"`
			}{
				Status: "Bad Gateway",
				Error:  err.Error(),
			})
		}
		return c.JSON(http.StatusCreated, u)
	})

	e.GET("/users/:id", func(c echo.Context) error {
		sid := c.Param("id")
		id, err := strconv.ParseInt(sid, 10, 64)
		us, err := client.User.
			Query().
			Where(user.IDEQ(int(id))). // Add Where
			Only(c.Request().Context())

		if !errors.Is(err, nil) {
			log.Fatalf("Error: failed quering users %v\n", err)
		}

		return c.JSON(http.StatusOK, us)
	})

	srv := handler.NewDefaultServer(graph.NewSchema(client))
	{
		e.POST("/query", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		e.GET("/playground", func(c echo.Context) error {
			playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_URL")))
}
