package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func buildURL(city string) string {
	baseURL := "https://api.openweathermap.org/data/2.5/forecast?q="
	appID := "&appid=634ab40d1fd51482f9e8d2891c5e45a0"
	units := "&metric"
	url := baseURL + city + units + appID
	return url
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hello! Where did you want to check the weather?")
	fmt.Println("Type the name of the city, then press enter")
	scanner.Scan()
	city := scanner.Text()
	fmt.Println("Retrieving the weather for", city)
	url := buildURL(city)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("finished", body)
}
