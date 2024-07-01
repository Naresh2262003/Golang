package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	myurl := "https://example.com:8080/path/to/resource?key1=value&key2=value2"

	parseURL, err := url.Parse(myurl)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("Type of url %T\n", parseURL)

	fmt.Println("Scheme :", parseURL.Scheme)
	fmt.Println("Host :", parseURL.Host)
	fmt.Println("Path :", parseURL.Path)
	fmt.Println("Raw Query :", parseURL.RawQuery)

	parseURL.Path = "/newPath"
	parseURL.RawQuery = "username=iamnaresh"

	newUrl := parseURL.String()

	fmt.Println("new Url", newUrl)

}
