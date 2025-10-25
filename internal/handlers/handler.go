package handlers

import (
	"encoding/json"
	"fmt"
	"io"
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
		c.JSON(http.StatusInternalServerError, gin.H{"Database doesn't found": exists})
		return
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse db path": ok})
		return
	}

	result, err := storage.Get_ListSingle_SpyCat(*psql, Name.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to get single cat": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Single Agent:": result})
}

func CreateSpyCats_Handler(c *gin.Context) {
	var Cat *app.Cats

	if err := json.NewDecoder(c.Request.Body).Decode(&Cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error decode json dates:": err})
		return
	}

	IsValid, err := ValidBreed(Cat.Breed)
	if IsValid == false {
		c.JSON(http.StatusBadRequest, gin.H{"Breed is not valid": IsValid})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error while valid breed": err})
		return
	}

	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Database doesn't found": exists})
		return
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse db path": ok})
		return
	}

	name, err := storage.CreateCat(*psql, Cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cat", "err message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Agent was successfully created!": name})
}

func ValidBreed(breed string) (bool, error) {

	Breeds := make(map[int]string)
	valid := false

	if len(Breeds) == 0 {
		resp, err := http.Get("https://api.thecatapi.com/v1/breeds")
		if err != nil {
			return false, fmt.Errorf("Unable to get api", err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {

			return false, fmt.Errorf("Error to read body response", err)
		}

		var raw []struct {
			Name string
		}
		if err := json.Unmarshal(body, &raw); err != nil {

			return false, fmt.Errorf("Fail unmarshal breeds:", err)
		}

		for i, b := range raw {
			Breeds[i] = b.Name

		}

		for i, _ := range Breeds {
			if Breeds[i] == breed {

				valid = true
				break
			}

		}

	} else {

		for i, _ := range Breeds {
			if Breeds[i] == breed {
				return true, nil
			}

		}
	}

	return valid, nil

}

func Get_ListAllCats_Handler(c *gin.Context) {

	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Database doesn't found": exists})
		return
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse db path": ok})
		return
	}

	name, err := storage.Show_AllSpyCats(*psql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to get all cats": err})
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
		c.JSON(http.StatusInternalServerError, gin.H{"Database doesn't found": exists})
		return
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse db path": ok})
		return
	}

	err := storage.DeleteCat(*psql, Cat.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to remove cat", "err message": err.Error()})
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{"Database doesn't found": exists})
		return
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse db path": ok})
		return
	}

	err := storage.UpdateCat(*psql, Cat.Salary, Cat.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Failed update salary": err})
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{"Database doesn't found": exists})
		return
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse db path": ok})
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{"Database doesn't found": exists})
		return
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse db path": ok})
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{"Database doesn't found": exists})
		return
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse db path": ok})
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{"Database doesn't found": exists})
		return
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse db path": ok})
		return
	}

	err := storage.UpdateMission(*psql, &Missions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Failed update salary": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Complete state was updated in": Missions.MissionsId})
}

func Get_AllMissions_Handler(c *gin.Context) {

	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Database doesn't found": exists})
		return
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse db path": ok})
		return
	}
	missions, err := storage.Show_AllMissions(*psql)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Failed get all missions": err})
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{"Database doesn't found": exists})
		return
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse db path": ok})
		return
	}

	result, err := storage.Get_SingleMission(*psql, Mission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to get single cat": err})
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{"Database doesn't found": exists})
		return
	}
	psql, ok := db.(*storage.Storage)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse db path": ok})
		return
	}

	err := storage.AssignCat(*psql, Mission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to get single cat": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"": Mission.Cat_id})

}
