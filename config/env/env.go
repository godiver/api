package env

import (
	"os"
)

// Config has all environment variables
type env struct {
	API API
}

var Env = env{}

func ReadEnv() {
	Env.API.Rakuten.ID = os.Getenv("RAKUTEN_APP_ID")
}
