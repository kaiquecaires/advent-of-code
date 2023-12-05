package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type AlmanacMap struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func findRanges(start int, end int, almanac []AlmanacMap, debug bool) [][2]int {
	ranges := make([][2]int, 0)

	biggestNumber := almanac[0].SourceRangeStart
	smallestNumber := almanac[0].SourceRangeStart

	for _, item := range almanac {
		diff := item.DestinationRangeStart - item.SourceRangeStart

		maxSourceRangeStart := item.SourceRangeStart + item.RangeLength - 1

		if start >= item.SourceRangeStart && start <= maxSourceRangeStart && end <= maxSourceRangeStart {
			ranges = append(ranges, [2]int{start + diff, end + diff})
		}

		if start >= item.SourceRangeStart && start <= maxSourceRangeStart && end > maxSourceRangeStart {
			ranges = append(ranges, [2]int{start + diff, maxSourceRangeStart + diff})
			start = maxSourceRangeStart + 1
		}

		if start < item.SourceRangeStart && end >= item.SourceRangeStart && end <= maxSourceRangeStart {
			ranges = append(ranges, [2]int{item.SourceRangeStart + diff, end + diff})
			end = item.SourceRangeStart - 1
		}

		if maxSourceRangeStart > biggestNumber {
			biggestNumber = maxSourceRangeStart
		}

		if item.DestinationRangeStart < smallestNumber {
			smallestNumber = item.DestinationRangeStart
		}
	}

	if start > biggestNumber || start < smallestNumber {
		ranges = append(ranges, [2]int{start, end})
	}

	return ranges
}

func getSmallestLocation(locations [][2]int) int {
	smallestLocation := locations[0][0]

	for i := 0; i < len(locations); i++ {
		if smallestLocation > locations[i][0] {
			smallestLocation = locations[i][0]
		}
	}

	return smallestLocation
}

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_05/input.txt")
	rows := strings.Split(string(f), "\n")

	seedValues := strings.Fields(
		strings.Split(rows[0], "seeds: ")[1],
	)

	almanac := make(map[string][]AlmanacMap)
	currentAlmanacKey := ""

	for row := 1; row < len(rows); row++ {
		if strings.Trim(rows[row], "") == "" {
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

	smallestNumbers := make([]float64, 0)

	for i := 0; i < len(seedValues); i = i + 2 {
		start, _ := strconv.Atoi(seedValues[i])
		interval, _ := strconv.Atoi(seedValues[i+1])
		soilRanges := findRanges(start, start+interval-1, almanac["seed-to-soil"], false)

		for _, soilRange := range soilRanges {
			fertilizerRanges := findRanges(soilRange[0], soilRange[1], almanac["soil-to-fertilizer"], false)

			for _, fertilizerRange := range fertilizerRanges {
				waterRanges := findRanges(fertilizerRange[0], fertilizerRange[1], almanac["fertilizer-to-water"], false)

				for _, waterRange := range waterRanges {
					lightRanges := findRanges(waterRange[0], waterRange[1], almanac["water-to-light"], false)

					for _, lightRange := range lightRanges {
						temperatureRanges := findRanges(lightRange[0], lightRange[1], almanac["light-to-temperature"], false)

						for _, temperatureRange := range temperatureRanges {
							humidityRanges := findRanges(temperatureRange[0], temperatureRange[1], almanac["temperature-to-humidity"], false)

							for _, humidityRange := range humidityRanges {
								locationRanges := findRanges(humidityRange[0], humidityRange[1], almanac["humidity-to-location"], true)
								smallestNumbers = append(smallestNumbers, float64(getSmallestLocation(locationRanges)))
							}
						}
					}
				}
			}
		}
	}

	smallestNumber := smallestNumbers[0]

	for _, number := range smallestNumbers {
		smallestNumber = math.Min(smallestNumber, number)
	}

	fmt.Println(smallestNumber)
}
