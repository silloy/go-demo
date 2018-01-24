package main

import "fmt"

func main()  {
	seasins := []string{"spring", "summer", "autumn", "winter"}
	for ix, sea := range seasins {
		fmt.Printf("Season %d is %s\n", ix, sea)
	}

	var season string
	for _, season = range seasins{
		fmt.Printf("%s\n", season)
	}
}
