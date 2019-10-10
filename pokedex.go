package main

import "fmt"

//Pokedex isa directory of pokemon
type Pokedex struct {
	directory []Pokemon
}

//Add will insert a new pokemon into the directory.
func (p *Pokedex) Add(pokemon Pokemon) {
	p.directory = append(p.directory, pokemon)
}

//List prints out all pokemon in directory
func (p *Pokedex) List() {
	for _, pokemon := range p.directory {
		fmt.Println(pokemon.name, pokemon.caught)
	}
}
