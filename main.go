package main

/*
    Program for generating desktop notification about current weather.

    Set your API key and city id below and run the program.

    ** Icons used here are taken from https://openweathermap.org/weather-conditions **
*/

import (
    "fmt"
    "log"
    "time"
    "net/http"
    "encoding/json"
)

const API_KEY = OWM_API_KEY;
const CITY_ID = "5969782" // https://openweathermap.org/find
const URL = "http://api.openweathermap.org/data/2.5/weather?"



//makeRequest will fetch the http response of a given URL
func makeRequest(url string) *http.Response {
    client := &http.Client{Timeout: 3*time.Second}
    resp, err := client.Get(url)
    failOnError(err, "failed to GET url")
    return resp
}

//https://stackoverflow.com/questions/17156371/
//parseResponse will populate the WeatherData struct
func parseResponse(resp *http.Response) WeatherData  {
    data := &WeatherData{}
    err :=  json.NewDecoder(resp.Body).Decode(data)
    defer resp.Body.Close()
    failOnError(err, "failed to decode body")
    return *data
}


//callNotification will setup the command string and call it
func callNotification(data WeatherData) {
    cmd := initCommandString();
    cmd = setIcon(cmd, data.icon())
    cmd = setTitle(cmd, data.tempC())
    cmd = setBody(cmd, fmt.Sprintf("%s\nmin: %s  max: %s",data.desc(), data.minTempC(), data.maxTempC()))
    err := runShellCommand(cmd)
    failOnError(err,"failed cmd")
}

//failOnError used for simplifying error handling
func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}



func main() {

    var url string
    var data WeatherData;
    var resp *http.Response;

    url = fmt.Sprintf("%sid=%s&appid=%s", URL, CITY_ID, API_KEY)
    resp = makeRequest(url)
    data = parseResponse(resp)

    callNotification(data)
}
