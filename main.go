package main

import (
	"fmt"
	"log"
	"github.com/qwerty22121998/gotor-changeip/gotor-changeip"
)

func main() {
	client, _ := gotor_changeip.NewClient()
	fmt.Println(client.CurrentIP())
	err := client.Renew()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client.CurrentIP())
	client.Close()

}
