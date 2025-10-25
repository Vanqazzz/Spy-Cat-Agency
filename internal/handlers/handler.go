package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	storage "main.go/Storage"
	"main.go/internal/app"
)

func Get_SingleSpyCat_Handler(c *gin.Context) {
	var Name app.Cats
	if err := json.NewDecoder(c.Request.Body).Decode(&Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed ", "err message": err.Error()})
		return
	}

	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	fmt.Println(Name)

	result, err := storage.Get_ListSingle_SpyCat(*psql, Name.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get single cat", "err message": err.Error()})
	}

	c.JSON(http.StatusAccepted, gin.H{"Single Agent:": result})
}

func CreateSpyCats_Handler(c *gin.Context) {
	var Cat *app.Cats

	if err := json.NewDecoder(c.Request.Body).Decode(&Cat); err != nil {
		fmt.Errorf(err.Error(), err)
		return
	}

	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	name, err := storage.CreateCat(*psql, Cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cat", "err message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Agent was successfully created!": name})
}

func Get_ListAllCats_Handler(c *gin.Context) {

	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}

	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	name, err := storage.Show_AllSpyCats(*psql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all cats", "err message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": name})
}

func RemoveCat_Handler(c *gin.Context) {
	var Cat app.Cats

	if err := json.NewDecoder(c.Request.Body).Decode(&Cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON", "err message": err.Error()})
		return
	}

	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}

	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	err := storage.DeleteCat(*psql, Cat.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to remove cat", "err message": err.Error()})
	}

	c.JSON(http.StatusNoContent, gin.H{"Cat was succesfully removed with name:": Cat.Name})
}

func UpdateCat_Handler(c *gin.Context) {
	var Cat app.Cats

	if err := json.NewDecoder(c.Request.Body).Decode(&Cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON", "err message": err.Error()})
		return
	}

	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}

	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	err := storage.UpdateCat(*psql, Cat.Salary, Cat.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed update salary", "err message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"Salary was updated in": Cat.Name})
}

func CreateMission_Handler(c *gin.Context) {

	var Missions *app.Missions

	if err := json.NewDecoder(c.Request.Body).Decode(&Missions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON", "err message": err.Error()})
		return
	}

	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}

	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	err := storage.CreateMission(*psql, Missions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create mission", "err message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": Missions})
}

func CreateTarget_Handler(c *gin.Context) {
	var Target *app.Target

	if err := json.NewDecoder(c.Request.Body).Decode(&Target); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON", "err message": err.Error()})
		return
	}

	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}

	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	err := storage.CreateTarget(*psql, Target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create target", "err message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Target was created": Target.Target_name})

}

func DeleteMission_Handler(c *gin.Context) {
	var Missions *app.Missions

	if err := json.NewDecoder(c.Request.Body).Decode(&Missions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON", "err message": err.Error()})
		return
	}

	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}

	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	err := storage.DeleteMission(*psql, Missions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete mission", "err message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"Mission was deleted!": Missions})

}

func UpdateMission_Handler(c *gin.Context) {

	var Missions app.Missions

	if err := json.NewDecoder(c.Request.Body).Decode(&Missions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON", "err message": err.Error()})
		return
	}

	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}

	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	err := storage.UpdateMission(*psql, &Missions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed update salary", "err message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"Complete state was updated in": Missions.MissionsId})
}

func Get_AllMissions_Handler(c *gin.Context) {

	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}

	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	missions, err := storage.Show_AllMissions(*psql)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed get all missions", "err message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"All missions:": missions})
}

func Get_SingleMission_Handler(c *gin.Context) {
	var Mission app.Missions
	if err := json.NewDecoder(c.Request.Body).Decode(&Mission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed ", "err message": err.Error()})
		return
	}

	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	result, err := storage.Get_SingleMission(*psql, Mission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get single cat", "err message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"Single Agent:": result})
}

func AssignCat(c *gin.Context) {

	var Mission app.Missions
	if err := json.NewDecoder(c.Request.Body).Decode(&Mission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed ", "err message": err.Error()})
		return
	}
	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	err := storage.AssignCat(*psql, Mission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get single cat", "err message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"": Mission.Cat_id})

}
