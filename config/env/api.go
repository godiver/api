package env

type rakuten struct {
	ID     string
	Secret string
}

// API struct has environment variables for API
type API struct {
	Rakuten rakuten
}
