package pokeapi

import "fmt"

type Resource struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

func (r *Resource) printNames() {
	for _, result := range r.Results {
		fmt.Println(result.Name)
	}
}
