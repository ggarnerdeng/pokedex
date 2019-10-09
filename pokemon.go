package main

import "fmt"

//Pokemon is a data struct representing each poke we caught or seen
type Pokemon struct {
	id     int
	name   string
	caught bool
}

//Create constructs a pokemon
func (p *Pokemon) Create(id int, name string, caught bool, pokedex []Pokemon) {
	newPokemon := Pokemon{
		id:     id,
		name:   name,
		caught: caught,
	}
	pokedex = append(pokedex, newPokemon)
}

//Read prints pokedex
func (p *Pokemon) Read(pokedex []Pokemon) {
	for _, p := range pokedex {
		fmt.Println(p.name, p.caught)
	}
}
