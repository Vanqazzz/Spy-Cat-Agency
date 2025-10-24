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
	r := gin.Default()
	db, err := storage.New(storagePath)
	if err != nil {
		slog.Error("Error while connecting to DB", err)
	}

	r.Use(Db_Middleware(db))
	r.POST("/createcat", handlers.CreateSpyCats_Handler)
	r.GET("/getasinglecat", handlers.SingleSpyCat_Handler)
	r.GET("/listallcats", handlers.ListAllCats_Handler)
	r.DELETE("/delete", handlers.RemoveCat_Handler)
	r.PUT("/update", handlers.UpdateCat_Handler)

	r.Run()

}

/* func ListSpyCats_Handler(c *gin.Context) {
	var Name Cats
	if err := json.NewDecoder(c.Request.Body).Decode(&Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed ", "err message": err.Error()})
		return
	}

	storagePath := os.Getenv("DB_PATH")

	storagew, err := storage.New(storagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to db", "err message": err.Error()})

	}

	fmt.Println(Name)

	result, err := storage.ListSingle_SpyCat(*storagew, Name.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select request", "err message": err.Error()})
	}

	c.JSON(http.StatusAccepted, gin.H{"data": result.Name})
}

func CreateSpyCats_Handler(c *gin.Context) {
	var Cat Cats
	if err := json.NewDecoder(c.Request.Body).Decode(&Cat); err != nil {
		fmt.Errorf(err.Error(), err)
		return
	}

	storagePath := os.Getenv("DB_PATH")

	storagew, err := storage.New(storagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to db", "err message": err.Error()})
		return
	}

	name, err := storage.CreateCat(*storagew, Cat.Name, Cat.YearsOfExperience, Cat.Breed, Cat.Salary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cat", "err message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": name})
}
*/
/* {
    "Name": "Whiskers",
    "YearsOfExperience": 5,
    "Breed": "Persian",
    "Salary": 1000
} */
