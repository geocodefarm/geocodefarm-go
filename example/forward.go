package main

import (
	"fmt"
	"log"

	"github.com/geocodefarm/geocodefarm-go/geocodefarm"
)

func main() {
	client := geocodefarm.NewClient("YOUR_API_KEY_HERE")

	resp, err := client.Forward("1600 Amphitheatre Parkway, Mountain View, CA")
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
