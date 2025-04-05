# Geocode.Farm Go SDK

A lightweight Go client for interacting with the [Geocode.Farm API](https://geocode.farm/).

This SDK allows you to perform forward and reverse geocoding using the Geocode.Farm API. It has no external dependencies, making it simple and fast.

## Installation

To install the Geocode.Farm Go SDK, use Go modules:

```bash
go get github.com/geocodefarm/geocodefarm-go
```

## Usage

### Forward Geocoding

To convert an address into latitude and longitude:

```go
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
```

### Reverse Geocoding

To convert latitude and longitude into an address:

```go
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
```

## Example Requests

- **Forward Geocoding**: Converts an address (e.g., `1600 Amphitheatre Parkway, Mountain View, CA`) into latitude and longitude.
- **Reverse Geocoding**: Converts latitude and longitude (e.g., `37.4221, -122.0841`) into a full address.

## License

This repository is licensed under the [Unlicense](https://unlicense.org/). See the LICENSE file for more information.
