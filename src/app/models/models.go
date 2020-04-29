package models

const (
	Small string = "small"
	Big string = "big"
)

// Bollard = Amarra
type Bollard struct {
	Type string `json:"type"`
	Free bool `json:"free"`
}

// Boats = Barcos
type Boats struct {
	Type string `json:"type"`
	Name string `json:"name"`
}
