package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/mtslzr/pokeapi-go"
)

type Pokemon struct {
	Name string
	Image string
	Id int
	Types []string
	Abilities []string
}

type ErrorMessage struct {
	Message error
}

func main() {
	r := gin.Default()
	r.GET("/", func (c *gin.Context)  {
		tmpl := template.Must(template.ParseFiles("./index.html"))
		bulbasaur,bulbasaurErr := createPokemon("bulbasaur")
		charmander, charmanderErr := createPokemon("charmander")
		squirtle, squirtleErr := createPokemon("squirtle")

		if bulbasaurErr != nil {
			errorMessage := ErrorMessage{
				Message: bulbasaurErr,
			}
			tmpl := template.Must(template.ParseFiles("./index.html"))
			tmpl.ExecuteTemplate(c.Writer,"error-toast", errorMessage)
			return
		}
		if charmanderErr != nil {
			errorMessage := ErrorMessage{
				Message: charmanderErr,
			}
			tmpl := template.Must(template.ParseFiles("./index.html"))
			tmpl.ExecuteTemplate(c.Writer,"error-toast", errorMessage)
			return
		}
		if squirtleErr != nil {
			errorMessage := ErrorMessage{
				Message: squirtleErr,
			}
			tmpl := template.Must(template.ParseFiles("./index.html"))
			tmpl.ExecuteTemplate(c.Writer,"error-toast", errorMessage)
			return
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
		c.JSON(http.StatusOK, gin.H{
			"message": "Everything is up and running!",
		})
	})
	r.POST("/pokemon", getPokemon)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getPokemon(c *gin.Context) {
	pokemonName := c.PostForm("pokemonName")
	pokemon, err := createPokemon(pokemonName)

if err != nil {
	log.Print(err)
	errorMessage := ErrorMessage{
		Message: err,
	}
	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.ExecuteTemplate(c.Writer,"error-toast", errorMessage)
	return
}	
	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.ExecuteTemplate(c.Writer, "pokemon-card", pokemon)
}

func createPokemon (pokemonName string) (Pokemon, error) {

	if pokemonName == "" {
		return Pokemon{}, errors.New("empty Pokemon Name received")
	}
	pokemonResponse, _ := pokeapi.Pokemon(strings.ToLower(pokemonName))
	if pokemonResponse.Name == "" {
		return Pokemon{}, errors.New("pokemon was not found")
	}
	pokemon := Pokemon{
		Name: pokemonResponse.Name,
		Image: pokemonResponse.Sprites.FrontDefault,
		Id: pokemonResponse.ID,
	}

	for i := 0; i < len(pokemonResponse.Types); i++ {
		pokemon.Types = append(pokemon.Types, pokemonResponse.Types[i].Type.Name)
	}

	for i := 0; i < len(pokemonResponse.Abilities); i++ {
		pokemon.Abilities = append(pokemon.Abilities, pokemonResponse.Abilities[i].Ability.Name)
	}

	return pokemon, nil
}