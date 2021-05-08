package main

import (
	"net/http"

  "github.com/fdaines/github-actions-example/model"
  "github.com/fdaines/github-actions-example/usecase"
	"github.com/labstack/echo/v4"
)
const PORT = ":9876"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
    data := usecase.RequestWorldBankData()
    
		return c.JSON(http.StatusOK, &ResponseDTO{
      Countries: data,
    })
	})

	e.Logger.Fatal(e.Start(PORT))
}

type ResponseDTO struct {
  Countries []model.CountryData `json:"countries"`
}
