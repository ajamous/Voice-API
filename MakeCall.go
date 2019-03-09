package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	var data = []byte(`{From=%2b19542405000&To=73297%2319541235001&Url=https://cloud.restcomm.com/restcomm-rvd/services/apps/AP1c9836a2f3d04e72a74f1e277d232e42/controller}`)
	req, err := http.NewRequest("POST", "https://ACcc38b3e0fcf3d2654cd73e210f542233ce:13dd0e03b368f49586428e7a5a3sc49a@cloud.restcomm.com/restcomm/2012-04-24/Accounts/ACcc38b3e0fcf3d2654cd73e210f5477ce/Calls.json", data)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
