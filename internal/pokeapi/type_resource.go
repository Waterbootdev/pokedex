package pokeapi

type Resource struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}
