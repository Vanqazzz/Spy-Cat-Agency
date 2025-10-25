package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	storage "main.go/Storage"
	"main.go/internal/handlers"
)

var (
	storagePath = os.Getenv("DB_PATH")
)

func Db_Middleware(db *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
func main() {

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
