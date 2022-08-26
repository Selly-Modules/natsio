package model

// LocationRequestPayload ...
type LocationRequestPayload struct {
	Province int `json:"province"`
	District int `json:"district"`
	Ward     int `json:"ward"`
}
