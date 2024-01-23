package main

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price  float64  `json:"price"`
}

var albums =[]album{{ID: "1",Title: "Blue Train",Artist: "Lemon",Price: 30.40},
{ID: "1",Title: "Last last",Artist: "Lemonr",Price: 40.40},
{ID: "1",Title: "Real Life",Artist: "Lemonrr",Price: 50.40},
{ID: "1",Title: "Ub40",Artist: "Lemonrrr",Price: 60.40},
{ID: "1",Title: "MoonNight",Artist: "Lemonrrrr",Price: 70.40},
{ID: "1",Title: "Halloed Jones",Artist: "Lemonrrrr",Price: 80.40},
{ID: "1",Title: "Sprinter",Artist: "Lemonrrrrrrr",Price: 90.40},
{ID: "1",Title: "Gollowed Tingest",Artist: "Lemonrrrrrr",Price: 130.40},
{ID: "1",Title: "degoes Denolin",Artist: "Lemonrrrrrrrr",Price: 140.40}}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK,albums)
}

func postAlbums(c *gin.Context){
	var newAlbum album
	if  err := c.BindJSON(&newAlbum); err !=nil{
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated,newAlbum)

}

func main(){
	router:= gin.Default()
	router.GET("/albums",getAlbums)
	router.POST("/albums",postAlbums)

	router.Run("localhost:8080")
}