package models

type NetworkFilter struct {
	Mac string `json:"mac"`
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
	DeviceId string `json:"device_id"`
}