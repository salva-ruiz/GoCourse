package main

import (
	"net/http"
	"fmt"
	"context"
)

func main() {
	sites := []string {
		"https://www.google.com",
		"https://drive.google.com",
		//"https://maps.google.com",
		"https://500.google.com",
		"https://hangouts.google.com",
	}

	result := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())

        for index, site := range sites {
		fmt.Print(index + 1, ". ", site)

		go check_site(ctx, result, site)
		status := <-result

		if status == false {
			fmt.Println("\nAborted!")
			break
		}
	}

	cancel()
}

func check_site (ctx context.Context, result chan<- bool, site string) {
	res, err := http.Get(site)
	status := (err == nil)

	if status {
		fmt.Println(" ->", res.Status)
	} else {
		fmt.Println(" ->", err)
	}

	select {
		case result <- status:
			return
	}
}
