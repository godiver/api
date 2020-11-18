package env

type rakuten struct {
	ID     string
	Secret string
}

type youtube struct {
	Key string
}

// API struct has environment variables for API
type API struct {
	Rakuten rakuten
	Youtube youtube
}
