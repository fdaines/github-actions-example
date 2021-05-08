package repository

import (
  "io/ioutil"
  "net/http"
)

func RequestToWorldBankService() ([]byte, error) {
  res, err := http.Get("http://api.worldbank.org/v2/country?format=json&per_page=100&region=LCN")
  if err != nil {
    return nil, err
  }
  data, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }
  res.Body.Close()

  return data, nil
}
