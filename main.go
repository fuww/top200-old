package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	POLYGON_API_KEY := os.Getenv("POLYGON_API_KEY")

	tickers := []string{
		"NKE",
		"TJX",
		"VFC",
		// "COH",
		"GPS",
		// "LB",
		// "TIF",
		"PVH",
		// "KORS",
		"M",
		"RL",
		"JWN",
		"HBI",
		"UA",
		"CRI",
		"FL",
		"COLM",
		"LULU",
		"URBN",
		"AEO",
		"DKS",
	}

	date := time.Date(2023, 11, 01, 0, 0, 0, 0, time.Local)

	// init client
	c := polygon.New(POLYGON_API_KEY)

	// Iterate over the slice using a for loop
	for _, ticker := range tickers {
		params := models.GetTickerDetailsParams{
			Ticker: ticker,
		}.WithDate(models.Date(date))

		// make request
		r, err := c.GetTickerDetails(context.Background(), params)

		if err != nil {
			log.Fatal(err)
		}

		log.Print(r) // do something with the result
		// log.Print(r.Results.Ticker)

		filename := r.Results.Ticker + "_" + date.Format("2006-01-02") + "_" + ".json"

		f, e := os.Create("data/" + filename)
		if e != nil {
			panic(e)
		}
		defer f.Close()

		jsonByte, _ := json.Marshal(r)
		reader := bytes.NewReader(jsonByte)
		io.Copy(f, reader)
	}

}
