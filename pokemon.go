package main

//Pokemon is a data struct representing each poke we caught or seen
type Pokemon struct {
	id     int
	name   string
	caught bool
}

//Create constructs a pokemon
func Create(id int, name string, caught bool, pokedex []Pokemon) {
	newPokemon := Pokemon{
		id:     id,
		name:   name,
		caught: caught,
	}
	pokedex = append(pokedex, newPokemon)
}

//NewPokemon du
func NewPokemon(id int, name string, caught bool) Pokemon {
	return Pokemon{
		id:     id,
		name:   name,
		caught: caught,
	}
}
