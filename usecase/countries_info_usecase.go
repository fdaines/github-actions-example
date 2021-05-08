package usecase

import (
  "encoding/json"
  "log"
  "github.com/fdaines/github-actions-example/model"
  "github.com/fdaines/github-actions-example/repository"
  "github.com/fdaines/github-actions-example/translator"
)

func RequestWorldBankData() []model.CountryData {
  responseBody, err := repository.RequestToWorldBankService()
  if err != nil {
    log.Fatal("Cannot access world bank service")
  }
  responseObject := []interface{}{}
  err = json.Unmarshal(responseBody, &responseObject)
  if err != nil {
    log.Fatal("Cannot access world bank service XXX")
	}
  return translator.TranslateToCountryEntries(responseObject[1].([]interface{}))
}
