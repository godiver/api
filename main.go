package main

import (
	"app/config/env"
	"app/router"
)

func main() {
	env.ReadEnv()
	router.Router()
}
