package main

import "fmt"

// CSICResponse ...
type CSICResponse struct {
	Descriptions map[string]string `json:"descriptions"`
	Soles        []sol             `json:"soles"`
}

type sol struct {
	ID            string `json:"id"`
	Date          string `json:"terrestrial_date"`
	Sol           string `json:"sol"`
	LS            string `json:"ls"`
	Season        string `json:"season"`
	TempMin       string `json:"min_temp"`
	TempMax       string `json:"max_temp"`
	Pressure      string `json:"pressure"`
	PressureStr   string `json:"pressure_string"`
	Humidity      string `json:"abs_humidity"`
	WindSpeed     string `json:"wind_speed"`
	WindDirection string `json:"wind_direction"`
	Sunrise       string `json:"sunrise"`
	Sunset        string `json:"sunset"`
	AtmoOpacity   string `json:"atmo_opacity"`
	UVIndex       string `json:"local_uv_irradiance_index"`
	TempGroundMin string `json:"min_gts_temp"`
	TempGroundMax string `json:"max_gts_temp"`
}

type fSol struct {
	ID            uint64
	Date          string
	Sol           uint64
	SolarLon      int64
	Month         uint64
	TempMin       int64
	TempMax       int64
	Pressure      int64
	PressureStr   string
	Humidity      string
	WindSpeed     string
	WindDirection string
	Sunrise       string
	Sunset        string
	AtmoOpacity   string
	UVIndex       string
	TempGroundMin int64
	TempGroundMax int64
}

func (f fSol) String() string {
	return fmt.Sprintf(
		"fSol: {\n  ID: %d\n  Date: %s\n  Sol: %d\n  Month: %d\n  Temp: (%d, %d)\n}",
		f.ID, f.Date, f.Sol, f.Month, f.TempMin, f.TempMax,
	)
}
