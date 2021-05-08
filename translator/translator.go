package translator

import (
  "strconv"
  "github.com/fdaines/github-actions-example/model"
)

func TranslateToCountryEntries(items []interface{}) []model.CountryData {
  var countryEntries []model.CountryData

  for _, v := range items {
    entry := v.(map[string]interface{})
    lat, _ := strconv.ParseFloat(entry["latitude"].(string), 8)
    lon, _ := strconv.ParseFloat(entry["longitude"].(string), 8)

    countryEntries = append(countryEntries, model.CountryData{
      Code: entry["iso2Code"].(string),
      Name: entry["name"].(string),
      CapitalCity: entry["capitalCity"].(string),
      Latitude: lat,
      Longitude: lon,
    })
  }

  return countryEntries
}
