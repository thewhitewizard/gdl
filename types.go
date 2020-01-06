package gdl


// Parking
type Parking struct {

	ID string `json:"Parking_schema:identifier"`
	Name string `json:"Parking_schema:name"`
	UpdatedTime  string `json:"dct:date"`
	Closed bool `json:"ferme"`
	AvailableStandardSpaces int `json:"mv:currentValue"`
}
