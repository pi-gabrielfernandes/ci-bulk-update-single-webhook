package gympass

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func UpdateSingleWebhook(c *http.Client, r UpdateWebhookRequestDTO) (success bool) {
	resource := "partners/{uuid}/webhooks"
	uri, err := url.JoinPath(os.Getenv("SETUP_API_BASE_URI"), strings.Replace(resource, "{uuid}", r.PartnerTagusUUID, 1))
	if err != nil {
		log.Fatal(err)
	}

	payload, err := json.Marshal(UpdateWebhookRequest{ Webhook: r.Webhook })
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("PUT", uri, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.AccessToken))
	req.Header.Set("Content-type", "application/json")
	
	fmt.Printf("[debug] update webhook request: %v\n", req)

	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("[debug] update webhook response: %v\n", res)
	
	return res.StatusCode == http.StatusNoContent
}

type UpdateWebhookRequestDTO struct {
	AccessToken string
	PartnerTagusUUID string
	Webhook  WebhookDetails
}

type UpdateWebhookRequest struct {
	Webhook  WebhookDetails `json:"webhook"`
}

type WebhookDetails struct{
	Event string `json:"event"`
	Secret string `json:"secret"`
	Url string `json:"url"`
	PersonalData bool `json:"personal_data"`
}
