package gympass

import (
	//"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Authorize(c *http.Client) string {
	uri, err:= url.Parse(os.Getenv("AUTH_API_BASE_URI"))
	payload := url.Values{}
	payload.Set("client_id", os.Getenv("GP_CLIENT_ID"))
	payload.Set("grant_type", os.Getenv("GP_GRANT_TYPE"))
	payload.Set("username", os.Getenv("GP_USERNAME"))
	payload.Set("password", os.Getenv("GP_PASSWORD"))
	
	req, err := http.NewRequest("POST", uri.String(), strings.NewReader(payload.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)

	var result AuthorizeResponse
	json.Unmarshal(resBody, &result)

	return result.AccessToken
}

type AuthorizeResponse struct {
	AccessToken string `json:"access_token"`
}

