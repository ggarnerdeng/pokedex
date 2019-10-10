package main

//var pokedex []Pokemon
import (
	"flag"
	"net/http"
	_ "net/http"
	"os"
	"strconv"
)

var pokedex Pokedex

func addPokemon(res http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(req.FormValue("id"))
	name := req.FormValue("name")
	var caught bool
	//if req.FormValue("caught") = "true" {
	//	caught = true
	//}
	//caught := req.FormValue("caught")
	httpPokemon := NewPokemon(id, name, caught)
	pokedex.Add(httpPokemon)
	pokedex.List()
}
func main() {
	pokedex := Pokedex{}
	bulbasaur := NewPokemon(1, "Bulbasaur", true)
	pokedex.Add(bulbasaur)

	httpport := flag.String("http", ":8080", "port value")
	flag.Parse()
	http.HandleFunc("/pokemon", addPokemon)
	http.ListenAndServe(*httpport, nil)

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
