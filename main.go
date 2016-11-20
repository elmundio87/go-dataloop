package main

import "fmt"

import (

	"github.com/elmundio87/go-dataloop/dataloop"
	"net/http"
	"os"
)

func main() {
	org := os.Args[1]
	key := os.Args[2]
	
  client := dataloop.NewClient(&http.Client{}, org, key)
  
  success, resp, err := client.Accounts.Create("wibble")
  
  fmt.Println(success)
  fmt.Println(resp)
  
  
	if err != nil {
		fmt.Println(err)
		return
	}
  
}
