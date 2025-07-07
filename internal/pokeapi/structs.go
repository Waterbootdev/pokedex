package pokeapi

type Resource struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Result    `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ResourcePage struct {
	Resource      Resource
	Current_index int
}
