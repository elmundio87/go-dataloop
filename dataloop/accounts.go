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

type Account struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Agent_keys []string `json:"agent_keys"`
}

type AccountList []Account

type AccountRoleSummary struct {
    Account string `json:"account"`
    Org string `json:"org"`
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

// Update an existing account name on Dataloop.
func (o *AccountService) Update(accountName string, newAccountName string) (*SuccessCreate, *http.Response, error) {

	success := new(SuccessCreate)

  requestBody := &AccountRequest{
    Name: newAccountName,
  }

	resp, err := o.sling.New().Put("accounts/" + accountName ).BodyJSON(requestBody).Receive(success, nil)

	return success, resp, err
}

// Update an existing account name on Dataloop.
func (o *AccountService) Patch(accountName string, newAccountName string) (*SuccessCreate, *http.Response, error) {

	success := new(SuccessCreate)

  requestBody := &AccountRequest{
    Name: newAccountName,
  }

	resp, err := o.sling.New().Patch("accounts/" + accountName ).BodyJSON(requestBody).Receive(success, nil)

	return success, resp, err
}

// Retrieve an account list from Dataloop.
func (o *AccountService) List() (*AccountList, *http.Response, error) {

  var dataloopAccounts = new(AccountList)

	resp, err := o.sling.New().Get("accounts").Receive(dataloopAccounts, nil)

	return dataloopAccounts, resp, err
}

// Retrieve a single account from Dataloop.
func (o *AccountService) Retrieve(accountName string) (*Account, *http.Response, error) {

	var dataloopAccount = new(Account)

	resp, err := o.sling.New().Get("accounts/" + accountName).Receive(dataloopAccount, nil)

	return dataloopAccount, resp, err
}

// Retrieve an account role summary from Dataloop.
func (o *AccountService) RoleSummary(accountName string) (*AccountRoleSummary, *http.Response, error) {

	var dataloopAccountRoleSummary = new(AccountRoleSummary)

	resp, err := o.sling.New().Get("accounts/" + accountName + "/features").Receive(dataloopAccountRoleSummary, nil)

	return dataloopAccountRoleSummary, resp, err
}

// Delete an existing account from Dataloop.
func (o *AccountService) Delete(accountName string) (*http.Response, error) {

	resp, err := o.sling.New().Delete("accounts/" + accountName ).Receive(nil, nil)

	return resp, err
}
