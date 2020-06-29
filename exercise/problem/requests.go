package main

import (
	//"io"
	"net/http"
	//"os
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup

	sites := []string {
		"https://www.google.com",
		"https://drive.google.com",
		//"https://maps.google.com",
		"https://500.google.com",
		"https://hangouts.google.com",
	}

	for index, site := range sites {
		wg.Add(1)

		go func(index int, site string) {
			defer wg.Done()

			res, err := http.Get(site)

			if err == nil {
				//io.WriteString(os.Stdout, res.Status + "\n")
				fmt.Println(index, "-", site, "->", res.Status)
			} else {
				fmt.Println(index, "-", site, "->", err)
			}
		}(index, site)
	}

	wg.Wait()
}
