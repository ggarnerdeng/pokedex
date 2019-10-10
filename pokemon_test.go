package main

import "testing"

func TestNewPokemon(t *testing.T) {
	testPokemon := NewPokemon(1, "Bulbasaur", true)
	if testPokemon.name == "Bulbasaur" {
		t.Log("the test pokemon bulbasaur was propertly named")
	} else {
		t.Errorf("test pokemon not bulbasaur, but %s", testPokemon.name)
	}

}

//go test -v
