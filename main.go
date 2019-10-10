package main

//var pokedex []Pokemon
import (
	"os"
	"strconv"
)

func main() {
	pokedex := Pokedex{}
	bulbasaur := NewPokemon(1, "Bulbasaur", true)
	pokedex.Add(bulbasaur)

	if len(os.Args) > 1 {
		id, _ := strconv.Atoi(os.Args[1])
		name := os.Args[2]
		var caught bool
		if os.Args[3] == "true" {
			caught = true
		} else {
			caught = false
		}
		newP := NewPokemon(id, name, caught)
		pokedex.Add(newP)
	}
	pokedex.List()
}
