package main

import "base/config"

func main() {
	// init config
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	// init logger
	// init database
	// init echo

	// graceful shotdown
}
