package main

type config struct {
	LocationAreas        []string
	currentIndex         int
	init                 bool
	next                 bool
	limit                int
	length               int
	currentLocationAreas []string
}

func incrementIndex(index int, value int, length int) int {
	next := index + value
	if next >= length {
		next -= length
	}
	return next
}

func decrementIndex(index int, value int, length int) int {
	next := index - value
	if next < 0 {
		next = length + next
	}
	return next
}

func increment(config *config) {
	config.currentIndex = incrementIndex(config.currentIndex, 1, config.length)
}

func decrement(config *config) {
	config.currentIndex = decrementIndex(config.currentIndex, 1, config.length)
}

func nextLocationAreas(config *config) {
	current := config.currentLocationAreas

	if config.init {
		config.init = false
		config.next = true
		config.currentIndex = 0
	} else if config.next {
	} else {
		config.next = true
		config.currentIndex = incrementIndex(config.currentIndex, config.limit+1, config.length)
	}

	for i := 0; i < config.limit; i++ {
		current[i] = config.LocationAreas[config.currentIndex]
		increment(config)
	}

}

func prevLocationAreas(config *config) {
	current := config.currentLocationAreas

	if config.init {
		config.init = false
		config.next = false
		config.currentIndex = config.length - 1
	} else if config.next {
		config.next = false
		config.currentIndex = decrementIndex(config.currentIndex, config.limit+1, config.length)
	} else {

	}

	for i := 0; i < config.limit; i++ {
		current[config.limit-1-i] = config.LocationAreas[config.currentIndex]
		decrement(config)
	}

}
