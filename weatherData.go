package main

import "fmt"

//https://openweathermap.org/current
//https://openweathermap.org/weather-conditions

const ZeroKelvin = -273.15

type WeatherMeta struct {
	Main string  `json:"main"`
	Description string `json:"description"`
	Icon string `json:"icon"`
}

type Forecast struct {
	Temp float32
	Pressure float32
	Humidity float32
	TempMin float32 `json:"temp_min"`
	TempMax float32 `json:"temp_max"`
}

type WindCast struct {
	Speed float32  `json:"speed"` //meters per second
	Deg float32	 `json:"deg"` 	//meteorological degree
}

type CloudCast struct {
	All int `json:"all"` // cloud percent 0..100
}

type RainCast struct {
	OneHr float32 `json:"1h"`  //volume in mm
	ThreeHr float32 `json:"3h"`
}

type SnowCast struct {
	OneHr float32 `json:"1h"`  //volume in mm
	ThreeHr float32 `json:"3h"`
}

type RegionCast struct {
	Country string `json:"country"` //country code
	Sunrise uint64	`json:"sunrise"`
	Sunset uint64	`json:"sunset"`
}

type WeatherData struct {
	Meta []WeatherMeta `json:"weather"`

	Main Forecast `json:"main"`

	Wind WindCast `json:"wind"`
	
	Clouds CloudCast `json:"cloud"`
	
	Rain RainCast `json:"rain"`

	Snow RainCast `json:"snow"`

	Sys RegionCast `json:"sys"`
	
	Timestamp uint64 `json:"dt"`

	Name string `json:"name"` //city name

}


//get icon path (must not be relative)
func (w WeatherData) icon() string {
	icon_location := "icons/"
	icon_label := w.Meta[0].Icon
	icon_ext := ".png"

	return fmt.Sprintf("%s%s%s", icon_location, icon_label, icon_ext)
}

//desc is the description of the weather
func (w WeatherData) desc() string {
	return w.Meta[0].Description
}

//current temperature in Celsius 
func (w WeatherData) tempC() string {
	return fmt.Sprintf("%.2f °C", w.Main.Temp + ZeroKelvin)
}

func (w WeatherData) minTempC() string {
	return fmt.Sprintf("%.1f °C", w.Main.TempMin + ZeroKelvin)
}

func (w WeatherData) maxTempC() string {
	return fmt.Sprintf("%.1f °C", w.Main.TempMax + ZeroKelvin)
}