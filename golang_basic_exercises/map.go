package main

import "fmt"

func main() {
	elements := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println(elements)
	fmt.Println(elements["three"])

	website := map[string]map[string]string{
		"Google": map[string]string{
			"name": "Google",
			"type": "Search",
		},
		"youtube": map[string]string{
			"name": "Google",
			"type": "video",
		},
	}

	fmt.Println(website["Google"]["name"])
}
