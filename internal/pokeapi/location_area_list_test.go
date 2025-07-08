package pokeapi

import (
	"testing"
	"time"

	"github.com/Waterbootdev/pokedex/internal/pokecache"
)

func TestNextPreviousLocationAreas(t *testing.T) {
	client := NewClient(5*time.Second, pokecache.NewCache(5*time.Second))

	first, err := client.ListLocationAreas(locationAreasURL)

	if err != nil {
		t.Error(err)
		return
	}

	next, err := client.ListLocationAreas(*first.Next)

	if err != nil {
		t.Error(err)
		return
	}

	previous, err := client.ListLocationAreas(*next.Previous)

	if err != nil {
		t.Error(err)
		return
	}

	if *first.Next != *previous.Next {
		t.Errorf("")
		return
	}

}
