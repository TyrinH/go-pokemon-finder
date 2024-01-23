package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mtslzr/pokeapi-go"
)

func main() {
	r := gin.Default()
	r.GET("/", func (c *gin.Context)  {
		
	})
	r.GET("/ping", func(c *gin.Context)  {
		pokemon, _ := pokeapi.Pokemon("pikachu")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"pokemon": pokemon.Name,
		})
	})
	r.GET("/pokemon/:name", getPokemon)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getPokemon(c *gin.Context) {
	pokemonName := c.Param("name")
	pokemon, _ := pokeapi.Pokemon(pokemonName)

	if pokemon.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "The Pokemon requested was not found.",
		})
		return
	}
	c.JSON(http.StatusOK, pokemon)
}