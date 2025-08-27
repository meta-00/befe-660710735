package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

type CAT struct {
    ID       string  `json:"id"`
    Name     string  `json:"name"`
    Breed    string  `json:"breed"`
    Year     int     `json:"year"`
    Price    int  `json:"price"`
}

var cats = []CAT{
    {ID: "1", Name: "Oliver", Breed: "Bengal", Year: 2, Price: 450},
    {ID: "2", Name: "Luna", Breed: "Persian", Year: 2, Price: 420},
	{ID: "3", Name: "Bella", Breed: "Ragdoll", Year: 1, Price: 380},
	{ID: "4", Name: "Simba", Breed: "British Shorthair", Year: 1, Price: 320},
	{ID: "5", Name: "Daisy", Breed: "Scottish Fold", Year: 1, Price: 410},
}

func getCats(c *gin.Context){
	yearQuery := c.Query("year")

	if yearQuery != ""{
		filter := []CAT{}
		for _, cat := range cats {
			if fmt.Sprint(cat.Year) == yearQuery {
				filter = append(filter, cat)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, cats)
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context){
		c.JSON(200, gin.H{"message" : "healthy"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/cats", getCats)
	}

	r.Run(":8080")

}