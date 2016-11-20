package dataloop

import (
	"net/http"
	"github.com/dghubble/sling"
)

const DataloopAPI = "https://api.dataloop.io/v1"

// Client is a Dataloop client for making Dataloop API requests.
type Client struct {
	sling *sling.Sling

	// Different Dataloop Services
	Accounts *AccountService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, dataloopOrg, dataloopRestApiKey string) *Client {

	base := sling.New().Client(httpClient).Base(DataloopAPI + "/orgs/" + dataloopOrg + "/").Set("Authorization", "Bearer " + dataloopRestApiKey).Set("content-type", "multipart/form-data").Set("boundary", "---011000010111000001101001")

	return &Client{
		Accounts: NewAccountService(base.New()),
	}
}
