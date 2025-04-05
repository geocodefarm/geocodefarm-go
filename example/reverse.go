package main

import (
	"fmt"
	"log"

	"github.com/geocodefarm/geocodefarm-go/geocodefarm"
)

func main() {
	client := geocodefarm.NewClient("YOUR_API_KEY_HERE")

	resp, err := client.Reverse(37.4221, -122.0841)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if !resp.Success {
		log.Fatalf("API Error: %s", resp.Error)
	}

	fmt.Println("Latitude:", *resp.Lat)
	fmt.Println("Longitude:", *resp.Lon)
	fmt.Println("Accuracy:", *resp.Accuracy)
	fmt.Println("Full Address:", *resp.FullAddress)
}
