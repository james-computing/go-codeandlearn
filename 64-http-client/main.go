package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://bored-api.appbrewery.com/random", nil)
	if err != nil {
		fmt.Println("error creating request: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error doing request", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("status code:", resp.StatusCode)
		return
	}

	var boringRes BoringResponse
	err = json.NewDecoder(resp.Body).Decode(&boringRes)
	if err != nil {
		fmt.Println("error reading json response body:", err)
		return
	}

	fmt.Printf("%+v\n", boringRes)
}

type BoringResponse struct {
	Participants  int     `json:"participants"`
	Price         float64 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
	Accessibility string  `json:"accessibility"`
}
