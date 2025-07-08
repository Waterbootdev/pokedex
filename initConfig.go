package main

import (
	"time"

	"github.com/Waterbootdev/pokedex/internal/pokeapi"
)

func InitConfig() config {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	return config{pokeapiClient: pokeClient}
}
