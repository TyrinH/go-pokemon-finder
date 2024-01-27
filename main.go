package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mtslzr/pokeapi-go"
)

type Pokemon struct {
	Name string
	Image string
	Id int
}

type ErrorMessage struct {
	Message string
}

func main() {
	r := gin.Default()
	r.GET("/", func (c *gin.Context)  {
		tmpl := template.Must(template.ParseFiles("./index.html"))
		bulbasaurResponse, _ := pokeapi.Pokemon("bulbasaur")
		charmanderResponse, _ := pokeapi.Pokemon("charmander")
		squirtleResponse, _ := pokeapi.Pokemon("squirtle")

		bulbasaur := Pokemon{
			Name: bulbasaurResponse.Name,
			Image: bulbasaurResponse.Sprites.FrontDefault,
		}
		charmander := Pokemon{
			Name: charmanderResponse.Name,
			Image: charmanderResponse.Sprites.FrontDefault,

		}
		squirtle := Pokemon{
			Name: squirtleResponse.Name,
			Image: squirtleResponse.Sprites.FrontDefault,

		}

		starterPokemon := map[string][]Pokemon{
			"Pokemons": {
				bulbasaur,
				charmander,
				squirtle,
				
			},
		}
		tmpl.Execute(c.Writer, starterPokemon)
	})
	r.GET("/ping", func(c *gin.Context)  {
		pokemon, _ := pokeapi.Pokemon("pikachu")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"pokemon": pokemon.Name,
		})
	})
	r.POST("/pokemon", getPokemon)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getPokemon(c *gin.Context) {
	pokemonName := c.PostForm("pokemonName")

if pokemonName == "" {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "The requested Pokemon was not found",
	})
	return
}
	pokemonResponse, _ := pokeapi.Pokemon(pokemonName)
	pokemonIdString := strconv.Itoa(pokemonResponse.ID)
	log.Print("Pokemon Types:", pokemonResponse.Types)
	pokemonFormsResponse, _ := pokeapi.EvolutionChain(pokemonIdString)   

	log.Print(pokemonFormsResponse)

	if pokemonResponse.Name == "" {
		// log.Fatal("Pokemon was not found")
		errorMessage := ErrorMessage {
			Message: "Pokemon was not found",
		}
		tmpl := template.Must(template.ParseFiles("./index.html"))
		tmpl.ExecuteTemplate(c.Writer,"error-toast", errorMessage)

		return

	}
           
	pokemon := Pokemon{
		Name: pokemonResponse.Name,
		Image: pokemonResponse.Sprites.FrontDefault,
		Id: pokemonResponse.ID,

	}

	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.ExecuteTemplate(c.Writer, "pokemon-card", pokemon)
}