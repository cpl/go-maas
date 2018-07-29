package main

// CSICResponse ...
type CSICResponse struct {
	Descriptions map[string]string `json:"descriptions"`
	Soles        []sol             `json:"soles"`
}

type sol struct {
	ID            string
	Date          string
	Sol           string
	LS            string
	Season        string
	TempMin       string
	TempMax       string
	Preassure     string
	PressureStr   string
	Humidity      string
	WindSpeed     string
	WindDirection string
	Sunrise       string
	Sunset        string
	AtmoOpacity   string
	UVIndex       string
	TempGroundMin string
	TempGroundMax string
}
