package main

import (
	"bufio"
	"encoding/json"
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
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	responseBody := string(body)
	var data map[string]interface{}
	error := json.Unmarshal([]byte(responseBody), &data)
	if error != nil {
		panic(error)
	}
	d := data["main"].(map[string]interface{})
	// for k, v := range d {
	// 	fmt.Printf("key[%s] value[%s]\n", k, v)
	// }
	fmt.Println("Here's the current weather in", city)
	fmt.Println("High:", d["temp_max"], "'C")
	fmt.Println("Low:", d["temp_min"], "'C")
	fmt.Println("Humidity:", d["humidity"])
}
