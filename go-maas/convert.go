package main

import (
	"log"
	"strconv"
)

const invalid = "--"

func sol2fsol(s sol) fSol {
	var f fSol
	var err error

	f.AtmoOpacity = s.AtmoOpacity
	f.Date = s.Date
	f.Humidity = s.Humidity
	f.PressureStr = s.PressureStr
	f.UVIndex = s.UVIndex
	f.WindDirection = s.WindDirection
	f.WindSpeed = s.WindSpeed
	f.Sunrise = s.Sunrise
	f.Sunset = s.Sunset

	f.ID, err = strconv.ParseUint(s.ID, 10, 64)
	if err != nil {
		log.Fatal("ID ", err)
	}

	f.SolarLon, err = strconv.ParseInt(s.LS, 10, 64)
	if err != nil {
		log.Fatal("Solar Lon ", err)
	}

	f.Month, err = strconv.ParseUint(s.Season[6:], 10, 64)
	if err != nil {
		log.Fatal("Month ", err)
	}

	if s.Pressure != invalid {
		f.Pressure, err = strconv.ParseInt(s.Pressure, 10, 64)
		if err != nil {
			log.Fatal("Pressure ", err)
		}
	}

	f.Sol, err = strconv.ParseUint(s.Sol, 10, 64)
	if err != nil {
		log.Fatal("Sol ", err)
	}

	if s.Pressure != invalid {
		f.TempGroundMax, err = strconv.ParseInt(s.TempGroundMax, 10, 64)
		if err != nil {
			log.Fatal("Temp Ground Max ", err)
		}
	}

	if s.Pressure != invalid {
		f.TempGroundMin, err = strconv.ParseInt(s.TempGroundMin, 10, 64)
		if err != nil {
			log.Fatal("Temp Ground Min ", err)
		}
	}

	if s.Pressure != invalid {
		f.TempMax, err = strconv.ParseInt(s.TempMax, 10, 64)
		if err != nil {
			log.Fatal("Temp Max ", err)
		}
	}

	if s.Pressure != invalid {
		f.TempMin, err = strconv.ParseInt(s.TempMin, 10, 64)
		if err != nil {
			log.Fatal("Temp Min ", err)
		}
	}

	return f
}
