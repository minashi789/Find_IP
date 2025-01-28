package models

type Location struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	City    string  `json:"city"`
	Region  string  `json:"regionName"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type Response struct {
	IP string `json:"ip"`
}
