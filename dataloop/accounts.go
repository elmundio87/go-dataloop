package dataloop

import (

	"net/http"

	"github.com/dghubble/sling"
)

// https://documenter.getpostman.com/view/505524/dataloop_public_api/2FyccR#accounts

// AccountService provides methods for accessing Dataloop Account API endpoints.
type AccountService struct {
	sling *sling.Sling
}

type AccountRequest struct {
    Name     string   `json:"name"`
}

// NewAccountService returns a new AccountService
func NewAccountService(sling *sling.Sling) *AccountService {
	return &AccountService{
		sling: sling.Path(""),
	}
}

// SuccessCreate represents a Dataloop API Object Success response
type SuccessCreate struct {
	id string `json:"id"`
	name  string `json:"name"`
}

// SuccessUpdate represents a Dataloop API Object Update successful update response
type SuccessUpdate struct {
  id string `json:"id"`
	name  string `json:"name"`
}

// Create a new account on Dataloop.
func (o *AccountService) Create(accountName string) (*SuccessCreate, *http.Response, error) {

	success := new(SuccessCreate)

  requestBody := &AccountRequest{
    Name: accountName,
  }

	resp, err := o.sling.New().Post("accounts").BodyJSON(requestBody).Receive(success, nil)

	return success, resp, err
}
