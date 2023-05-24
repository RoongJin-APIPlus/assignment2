package graph

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/RoongJin/pokedex-graphql-sqlite/graph/model"
	"github.com/lib/pq"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Database struct {
	*sql.DB
}

type Resolver struct {
	DB Database
}

func (db Database) AddPokemon(name string, description string, category string, typeOf []string, abilities []string) (int64, error) {
	stmt := `INSERT INTO "pokedex"("name", "description", "category", "type", "abilities") values($1,$2,$3,$4,$5)`

	_, err := db.Exec(stmt, name, description, category, pq.Array(typeOf), pq.Array(abilities))
	if err != nil {
		return -1, err
	}

	rows, err := db.Query(`select * from "pokedex" where "name"=$1`, name)
	if err != nil {
		return -1, err
	}
	var d1 string
	var d2 string
	var d3 string
	var d4 string
	var d5 string
	var dummy string

	for rows.Next() {
		err = rows.Scan(&d1, &d2, &d3, &d4, &d5, &dummy)
		fmt.Println("Name: " + d1)
		fmt.Println("Description: " + d2)
		fmt.Println("Category: " + d3)
		fmt.Println("Type: " + d4)
		fmt.Println("Abilities: " + d5)
	}

	id_int64, _ := strconv.ParseInt(dummy, 10, 64)
	return id_int64, nil
}

func (db Database) UpdatePokemon(id int, name string, description string, category string, pokemonType []string, abilities []string) (bool, error) {
	stmt := `update "pokedex" set "name"=$1, "description"=$2, "category"=$3, "type"=$4, "abilities"=$5 where "id"=$6`

	_, err := db.Exec(stmt, name, description, category, pq.Array(pokemonType), pq.Array(abilities), id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (db Database) DeletePokemon(id int) (bool, error) {
	stmt := `delete from "pokedex" where "id"=$1`

	_, err := db.Exec(stmt, id)
	if err != nil {
		return false, err
	}

	return true, err
}

func (db Database) GetAllPokemons() ([]*model.Pokemon, error) {
	rows, err := db.Query(`select * from "pokedex"`)
	var pokeList []*model.Pokemon
	if err != nil {
		return pokeList, err
	}

	for rows.Next() {
		var name string
		var desc string
		var category string
		var types string
		var abilities string
		var dummy string
		err = rows.Scan(&name, &desc, &category, &types, &abilities, &dummy)
		fmt.Println("Name: " + name)
		fmt.Println("Description: " + desc)
		fmt.Println("Category: " + category)
		fmt.Println("Type: " + types)
		fmt.Println("Abilities: " + abilities)

		t := strings.Split(types, " ")
		a := strings.Split(abilities, " ")

		poke := model.Pokemon{
			ID:          dummy,
			Name:        name,
			Description: desc,
			Category:    category,
			Type:        t,
			Abilities:   a,
		}
		pokeList = append(pokeList, &poke)
	}

	defer rows.Close()

	return pokeList, nil
}

func (db Database) FindPokemonById(id int64) (model.Pokemon, error) {
	rows, err := db.Query(`select * from "pokedex" where "id"=$1`, id)
	if err != nil {
		return model.Pokemon{}, err
	}
	var name string
	var desc string
	var category string
	var types string
	var abilities string
	var dummy string

	for rows.Next() {
		err = rows.Scan(&name, &desc, &category, &types, &abilities, &dummy)
		fmt.Println("Name: " + name)
		fmt.Println("Description: " + desc)
		fmt.Println("Category: " + category)
		fmt.Println("Type: " + types)
		fmt.Println("Abilities: " + abilities)
	}

	dm := model.Pokemon{}
	if name == "" {
		return dm, fmt.Errorf("Pokemon with this ID does not exist!")
	}

	t := strings.Split(types, " ")
	a := strings.Split(abilities, " ")

	poke := model.Pokemon{
		ID:          dummy,
		Name:        name,
		Description: desc,
		Category:    category,
		Type:        t,
		Abilities:   a,
	}

	defer rows.Close()
	return poke, nil
}
