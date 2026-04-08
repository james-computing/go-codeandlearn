package main

import (
	"fmt"
	"net/http"
)

func main() {
	example1()
	example2()
}

type Result struct {
	resp *http.Response
	err  error
}

func checkIfExists(done <-chan struct{}, urls ...string) <-chan Result {
	results := make(chan Result)

	go func() {
		for _, url := range urls {
			select {
			case <-done:
				fmt.Println("Done")
				return
			default:
				resp, err := http.Get(url)
				results <- Result{resp: resp, err: err}
			}
		}
		close(results)
	}()

	return results
}

func example1() {
	done := make(chan struct{})

	results := checkIfExists(done, "https://duckduckgo.com", "https://localhost:300")

	for r := range results {
		if r.err != nil {
			fmt.Println(r.err)
		}
		if r.resp != nil {
			fmt.Println(r.resp.Status)
		}
	}

	close(done)
}

func checkIfExists2(done <-chan struct{}, urls ...string) (<-chan *http.Response, <-chan error) {
	responsec := make(chan *http.Response)
	errorc := make(chan error)

	go func() {
		for _, url := range urls {
			select {
			case <-done:
				fmt.Println("Done")
				return
			default:
				resp, err := http.Get(url)
				if err != nil {
					errorc <- err
				} else {
					responsec <- resp
				}
			}
		}
		close(responsec)
		close(errorc)
	}()

	return responsec, errorc
}

func example2() {
	done := make(chan struct{})

	responsec, errorc := checkIfExists2(done, "https://duckduckgo.com", "https://localhost:300")

	for range 2 {
		select {
		case response := <-responsec:
			fmt.Println(response.Status)
		case err := <-errorc:
			fmt.Println(err)
		}
	}

	close(done)
}
