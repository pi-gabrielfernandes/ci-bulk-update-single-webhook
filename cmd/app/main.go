package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/pi-gabrielfernandes/enable-personal-data-from-csv/internal/gympass"
	"github.com/pi-gabrielfernandes/enable-personal-data-from-csv/internal/utils"
)

var c = &http.Client{}
var token string

func main() {
	LoadEnv()
	token = gympass.Authorize(c)
	if token == "" {
		log.Fatal("unable to generate token")
	}
	utils.ProcessCsvFile(
		"input.csv",
		Process,
	)
}

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
}

func Process(row []string) error {
	fmt.Printf("[info] gym_id: %v\n", row[0])
	req := gympass.UpdateWebhookRequestDTO{	
		AccessToken: token,
		PartnerTagusUUID: row[1],
		Webhook: gympass.WebhookDetails{
			Event: "checkin",
			Url: row[2],
			Secret: row[3],
			PersonalData: true,
		},
	}
	res := gympass.UpdateSingleWebhook(c, req)
	fmt.Printf("result: %v\n", res)	
	return nil
}

