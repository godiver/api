package main

import (
	"src/github.com/godiver/api/config/env"
	"src/github.com/godiver/api/router"
)

func main() {
	env.ReadEnv()
	router.Router()
}
