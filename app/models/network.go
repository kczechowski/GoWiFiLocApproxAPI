package models

type Network struct {
	Mac string `json:"mac"`
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}
