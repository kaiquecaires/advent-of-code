package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AlmanacMap struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

type Seed struct {
	Value       int
	Soil        int
	Fertilizer  int
	Water       int
	Light       int
	Temperature int
	Humidity    int
	Location    int
}

func convertSeedToSoil(almanac map[string][]AlmanacMap, seed *Seed) {
	for _, currentAlmanac := range almanac["seed-to-soil"] {
		if seed.Value >= currentAlmanac.SourceRangeStart && seed.Value <= currentAlmanac.SourceRangeStart+currentAlmanac.RangeLength-1 {
			(*seed).Soil = currentAlmanac.DestinationRangeStart + (seed.Value - currentAlmanac.SourceRangeStart)
		}
	}
	if (*seed).Soil == 0 {
		(*seed).Soil = seed.Value
	}
}

func convertSoilToFertilizer(almanac map[string][]AlmanacMap, seed *Seed) {
	for _, currentAlmanac := range almanac["soil-to-fertilizer"] {
		if seed.Soil >= currentAlmanac.SourceRangeStart && seed.Soil <= currentAlmanac.SourceRangeStart+currentAlmanac.RangeLength-1 {
			(*seed).Fertilizer = currentAlmanac.DestinationRangeStart + (seed.Soil - currentAlmanac.SourceRangeStart)
		}
	}
	if (*seed).Fertilizer == 0 {
		(*seed).Fertilizer = seed.Soil
	}
}

func convertFertilizerToWater(almanac map[string][]AlmanacMap, seed *Seed) {
	for _, currentAlmanac := range almanac["fertilizer-to-water"] {
		if seed.Fertilizer >= currentAlmanac.SourceRangeStart && seed.Fertilizer <= currentAlmanac.SourceRangeStart+currentAlmanac.RangeLength-1 {
			(*seed).Water = currentAlmanac.DestinationRangeStart + (seed.Fertilizer - currentAlmanac.SourceRangeStart)
		}
	}
	if (*seed).Water == 0 {
		(*seed).Water = seed.Fertilizer
	}
}

func convertWaterToLight(almanac map[string][]AlmanacMap, seed *Seed) {
	for _, currentAlmanac := range almanac["water-to-light"] {
		if seed.Water >= currentAlmanac.SourceRangeStart && seed.Water <= currentAlmanac.SourceRangeStart+currentAlmanac.RangeLength-1 {
			(*seed).Light = currentAlmanac.DestinationRangeStart + (seed.Water - currentAlmanac.SourceRangeStart)
		}
	}
	if (*seed).Light == 0 {
		(*seed).Light = seed.Water
	}
}

func convertLightToTemperature(almanac map[string][]AlmanacMap, seed *Seed) {
	for _, currentAlmanac := range almanac["light-to-temperature"] {
		if seed.Light >= currentAlmanac.SourceRangeStart && seed.Light <= currentAlmanac.SourceRangeStart+currentAlmanac.RangeLength-1 {
			(*seed).Temperature = currentAlmanac.DestinationRangeStart + (seed.Light - currentAlmanac.SourceRangeStart)
		}
	}
	if (*seed).Temperature == 0 {
		(*seed).Temperature = seed.Light
	}
}

func convertTemperatureToHumidity(almanac map[string][]AlmanacMap, seed *Seed) {
	for _, currentAlmanac := range almanac["temperature-to-humidity"] {
		if seed.Temperature >= currentAlmanac.SourceRangeStart && seed.Temperature <= currentAlmanac.SourceRangeStart+currentAlmanac.RangeLength-1 {
			(*seed).Humidity = currentAlmanac.DestinationRangeStart + (seed.Temperature - currentAlmanac.SourceRangeStart)
		}
	}
	if (*seed).Humidity == 0 {
		(*seed).Humidity = seed.Temperature
	}
}

func convertHumidityToLocation(almanac map[string][]AlmanacMap, seed *Seed) {
	for _, currentAlmanac := range almanac["humidity-to-location"] {
		if seed.Humidity >= currentAlmanac.SourceRangeStart && seed.Humidity <= currentAlmanac.SourceRangeStart+currentAlmanac.RangeLength-1 {
			(*seed).Location = currentAlmanac.DestinationRangeStart + (seed.Humidity - currentAlmanac.SourceRangeStart)
		}
	}
	if (*seed).Location == 0 {
		(*seed).Location = seed.Humidity
	}
}

func convert(almanac map[string][]AlmanacMap, almanacKey string, seeds *[]Seed) {
	convertMap := map[string]func(almanac map[string][]AlmanacMap, seed *Seed){
		"seed-to-soil":            convertSeedToSoil,
		"soil-to-fertilizer":      convertSoilToFertilizer,
		"fertilizer-to-water":     convertFertilizerToWater,
		"water-to-light":          convertWaterToLight,
		"light-to-temperature":    convertLightToTemperature,
		"temperature-to-humidity": convertTemperatureToHumidity,
		"humidity-to-location":    convertHumidityToLocation,
	}

	for seedKey := range *seeds {
		convertMap[almanacKey](almanac, &(*seeds)[seedKey])
	}
}

func getNearestLocation(seeds *[]Seed) int {
	nearestLocation := (*seeds)[0].Location

	for i := 1; i < len(*seeds); i++ {
		if (*seeds)[i].Location < nearestLocation {
			nearestLocation = (*seeds)[i].Location
		}
	}

	return nearestLocation
}

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_05/input.txt")
	rows := strings.Split(string(f), "\n")

	seedValues := strings.Fields(
		strings.Split(rows[0], "seeds: ")[1],
	)

	seeds := make([]Seed, 0)

	for _, seedValue := range seedValues {
		value, _ := strconv.Atoi(seedValue)
		seeds = append(seeds, Seed{Value: value})
	}

	almanac := make(map[string][]AlmanacMap)
	currentAlmanacKey := ""

	for row := 1; row < len(rows); row++ {
		if strings.Trim(rows[row], "") == "" || row == len(rows)-1 {
			if len(almanac) > 0 {
				convert(almanac, currentAlmanacKey, &seeds)
			}

			currentAlmanacKey = ""
			continue
		}

		firstColumnString := strings.Split(rows[row], " ")[0]

		if _, ok := strconv.Atoi(firstColumnString); ok != nil {
			currentAlmanacKey = firstColumnString
			continue
		}

		if currentAlmanacKey != "" {
			numbers := strings.Split(rows[row], " ")

			destinationRangeStart, _ := strconv.Atoi(numbers[0])
			sourceRangeStart, _ := strconv.Atoi(numbers[1])
			rangeLength, _ := strconv.Atoi(numbers[2])

			almanac[currentAlmanacKey] = append(almanac[currentAlmanacKey], AlmanacMap{
				DestinationRangeStart: destinationRangeStart,
				SourceRangeStart:      sourceRangeStart,
				RangeLength:           rangeLength,
			})
		}
	}

	fmt.Println(getNearestLocation(&seeds))
}
