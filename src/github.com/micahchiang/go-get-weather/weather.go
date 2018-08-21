package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func buildURL(city string) string {
	baseURL := "http://api.openweathermap.org/data/2.5/weather?q="
	appID := "&appid=634ab40d1fd51482f9e8d2891c5e45a0"
	units := "&units=metric"
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
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(res)
	fmt.Println("finished", string(body))
}
