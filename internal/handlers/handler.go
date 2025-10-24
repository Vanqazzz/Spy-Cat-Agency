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

func SingleSpyCat_Handler(c *gin.Context) {
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

	result, err := storage.ListSingle_SpyCat(*psql, Name.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get single cat", "err message": err.Error()})
	}

	c.JSON(http.StatusAccepted, gin.H{"Single Agent:": result})
}

func CreateSpyCats_Handler(c *gin.Context) {
	var Cat app.Cats

	if err := json.NewDecoder(c.Request.Body).Decode(&Cat); err != nil {
		fmt.Errorf(err.Error(), err)
		return
	}

	/* resp, err := http.Get("https://api.thecatapi.com/v1/breeds")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to compare breed": err.Error()})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to read body response": err.Error()})
	}

	var breed map[string]interface{}
	if err := json.Unmarshal(body, &breed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to unmarsal": err.Error()})

	}

	fmt.Println(breed) */

	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	name, err := storage.CreateCat(*psql, Cat.Name, Cat.YearsOfExperience, Cat.Breed, Cat.Salary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cat", "err message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Agent was successfully created!": name})
}

func ListAllCats_Handler(c *gin.Context) {

	db, exists := c.Get("db")
	if !exists {
		slog.Error("Database doesn't found")
	}

	psql, ok := db.(*storage.Storage)
	if !ok {
		slog.Error("Error to parse database path")
	}

	name, err := storage.AllSpyCats(*psql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all cats", "err message": err.Error()})
		return
	}

	fmt.Println(name)

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

	c.JSON(http.StatusOK, gin.H{"Cat was succesfully removed with name:": Cat.Name})
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
