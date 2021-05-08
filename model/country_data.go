package model

type CountryData struct {
  Code string `json:"code"`
  Name string `json:"name"`
  CapitalCity string `json:"capital_city"`
  Latitude float64 `json:"latitude"`
  Longitude float64 `json:"longitude"`
}
