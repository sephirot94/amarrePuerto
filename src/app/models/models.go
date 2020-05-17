package models



const (
	Small string = "small"
	Big string = "big"
)

// Bollard = Amarra
type Bollard struct {
	Size string `json:"type"`
	Free bool `json:"free"`
}

// Boats = Barcos
type Boat struct {
	Size string `json:"type"`
	Name string `json:"name"`
}
