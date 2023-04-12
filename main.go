package main

import (
	"fmt"
	"os"

	"github.com/devshoe/abdm/pkg/abdm"
)

var (
	clientid     = os.Getenv("ABDM_CLIENT_ID")
	clientsecret = os.Getenv("ABDM_CLIENT_SECRET")
	client       *abdm.Client
)

func init() {
	if clientid == "" {
		panic("ABDM_CLIENT_ID not set")
	} else if clientsecret == "" {
		panic("ABDM_CLIENT_SECRET not set")
	} else {
		client = abdm.New(clientid, clientsecret)
	}
}

func main() {

	if err := client.Authenticate(); err != nil {
		panic(err)
	} else {
		fmt.Printf("%+v", client)
	}

}
