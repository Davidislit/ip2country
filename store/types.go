package store

type Location struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

type DB interface {
	Find(ip string) (*Location, error)
}
