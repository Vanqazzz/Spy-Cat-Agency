package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate"
	storage "main.go/Storage"
	"main.go/internal/handlers"
)

var (
	storagePath = os.Getenv("DB_PATH")
)

func Migrations() error {
	dbHost := os.Getenv("DB_PATH")
	migrationsPath := "file://./migrations"

	m, err := migrate.New(migrationsPath, storagePath)
	if err != nil {
		fmt.Println(dbHost)
		panic(fmt.Errorf("failed to create migrate instance: %w", err))
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(fmt.Errorf("failed to apply migrations: %w", err))
	}
	return nil
}

func Db_Middleware(db *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
func main() {

	Migrations()
	db, err := storage.New(storagePath)
	if err != nil {
		slog.Error("Error while connecting to DB", err)
	}

	r := gin.Default()

	r.Use(Db_Middleware(db))

	r.GET("/getasinglecat", handlers.Get_SingleSpyCat_Handler)
	r.GET("/listallcats", handlers.Get_ListAllCats_Handler)

	r.GET("/singlemission", handlers.Get_SingleMission_Handler)
	r.GET("/allmissions", handlers.Get_AllMissions_Handler)

	r.POST("/createmission", handlers.CreateMission_Handler)
	r.POST("/createcat", handlers.CreateSpyCats_Handler)
	r.POST("/createtarget", handlers.CreateTarget_Handler)

	r.PUT("/assigncat", handlers.AssignCat)

	r.DELETE("/delete", handlers.RemoveCat_Handler)
	r.DELETE("/deletemission", handlers.DeleteMission_Handler)

	r.PUT("/update", handlers.UpdateCat_Handler)
	r.PUT("/missionupdate", handlers.UpdateMission_Handler)

	r.Run()

}
