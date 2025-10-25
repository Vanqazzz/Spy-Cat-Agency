package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	storage "main.go/Storage"
	"main.go/internal/handlers"
)

var (
	storagePath = os.Getenv("Docker_DB_Path")
)

func Db_Middleware(db *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func Migrate(DBpath string) {

	m, err := migrate.New("file://./migrations",
		DBpath)
	if err != nil {
		fmt.Println(DBpath, err)
		panic("fail to apply migrations")

	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func main() {

	r := gin.Default()

	db, err := storage.New(storagePath)
	if err != nil {
		slog.Error("Error while connecting to DB", err)
	}
	Migrate(storagePath)

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
