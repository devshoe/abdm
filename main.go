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

// ABDM_CLIENT_ID="SBX_002660" ABDM_CLIENT_SECRET="4fd0ee05-ba91-4a4a-88e3-17d9e6fdd990"
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
		fmt.Println(client.RegisterHIU("devshoe", "hiu for second test", []string{"devshoe_2"}))
		fmt.Println(client.GetRegisteredServices())
	}

}
