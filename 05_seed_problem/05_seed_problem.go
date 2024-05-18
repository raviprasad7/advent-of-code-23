package seed_problem

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type InputData struct {
	Seeds                    []int
	SeedRanges               []SeedRange
	SeedToSoilMap            []EntityMapping
	SoilToFertilizerMap      []EntityMapping
	FertilizerToWaterMap     []EntityMapping
	WaterToLightMap          []EntityMapping
	LightToTemperatureMap    []EntityMapping
	TemperatureToHumidityMap []EntityMapping
	HumidityToLocationMap    []EntityMapping
}

type SeedRange struct {
	Start int
	End   int
}

type EntityMapping struct {
	Source int
	Delta  int
	Range  int
}

func Run() {
	var (
		inputData       InputData = InputData{}
		closestLocation int
	)

	file, err := os.Open("./05_seed_problem/sample_input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "seeds") {
			parts := strings.Split(line, ": ")
			seeds := strings.Split(parts[1], " ")

			for _, seed := range seeds {
				seedId, _ := strconv.Atoi(seed)
				inputData.Seeds = append(inputData.Seeds, seedId)
			}
		} else if strings.Contains(line, "map") {
			parseInputData(line, scanner, &inputData)
		}
	}

	// part 1
	closestLocation = findClosestLocation(inputData)
	fmt.Println("Part One - Lowest location number:", closestLocation)

	// part 2
	seedRanges := []SeedRange{}

	for i := 0; i < len(inputData.Seeds); i += 2 {
		seedRanges = append(seedRanges, SeedRange{
			Start: inputData.Seeds[i],
			End:   inputData.Seeds[i+1],
		})
	}
	inputData.SeedRanges = seedRanges
	// inputData.Seeds = []int{}

	// for _, seedRange := range seedRanges {
	// 	for i := 0; i < seedRange.End; i++ {
	// 		seedNo := seedRange.Start + i
	// 		inputData.Seeds = append(inputData.Seeds, seedNo)
	// 	}
	// }
	// fmt.Println("Seed ranges", seedRanges, inputData.Seeds)
	closestLocation = findClosestLocationInRange(inputData)
	fmt.Println("Part Two - Lowest location number with seed ranges:", closestLocation)
}

func parseInputData(currLine string, scanner *bufio.Scanner, inputData *InputData) {
	parts := strings.Split(currLine, " ")
	mapList := []EntityMapping{}

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			parts := strings.Split(line, " ")
			source, _ := strconv.Atoi(parts[1])
			destination, _ := strconv.Atoi(parts[0])
			sourceRange, _ := strconv.Atoi(parts[2])

			mapList = append(mapList, EntityMapping{
				Source: source,
				Delta:  destination - source,
				Range:  sourceRange,
			})
		} else {
			break
		}
	}

	switch parts[0] {
	case "seed-to-soil":
		inputData.SeedToSoilMap = mapList
	case "soil-to-fertilizer":
		inputData.SoilToFertilizerMap = mapList
	case "fertilizer-to-water":
		inputData.FertilizerToWaterMap = mapList
	case "water-to-light":
		inputData.WaterToLightMap = mapList
	case "light-to-temperature":
		inputData.LightToTemperatureMap = mapList
	case "temperature-to-humidity":
		inputData.TemperatureToHumidityMap = mapList
	case "humidity-to-location":
		inputData.HumidityToLocationMap = mapList
	}
}

func findClosestLocation(inputData InputData) int {
	closestLocation := math.MaxInt

	for _, seed := range inputData.Seeds {
		soil := findClosestMin(seed, inputData.SeedToSoilMap)
		fertilizer := findClosestMin(soil, inputData.SoilToFertilizerMap)
		water := findClosestMin(fertilizer, inputData.FertilizerToWaterMap)
		light := findClosestMin(water, inputData.WaterToLightMap)
		temperature := findClosestMin(light, inputData.LightToTemperatureMap)
		humidity := findClosestMin(temperature, inputData.TemperatureToHumidityMap)
		location := findClosestMin(humidity, inputData.HumidityToLocationMap)

		if location < closestLocation {
			closestLocation = location
		}

		// fmt.Println("Soil", soil, "Seed", seed, "Fertilizer", fertilizer, "Water", water, "Light", light, "Temperature", temperature, "Humidity", humidity, "Location", location)
	}

	return closestLocation
}

func findClosestLocationInRange(inputData InputData) int {
	closestLocation := math.MaxInt

	for i := 0; i < len(inputData.Seeds); i += 2 {
		fmt.Println("Seed range", inputData.SeedRanges)
		for j := inputData.Seeds[i]; j < inputData.Seeds[i+1]; j += 50000 {
			soil := findClosestMin(j, inputData.SeedToSoilMap)
			fertilizer := findClosestMin(soil, inputData.SoilToFertilizerMap)
			water := findClosestMin(fertilizer, inputData.FertilizerToWaterMap)
			light := findClosestMin(water, inputData.WaterToLightMap)
			temperature := findClosestMin(light, inputData.LightToTemperatureMap)
			humidity := findClosestMin(temperature, inputData.TemperatureToHumidityMap)
			location := findClosestMin(humidity, inputData.HumidityToLocationMap)

			if location < closestLocation {
				closestLocation = location
			}

			// fmt.Println("Soil", soil, "Seed", j, "Fertilizer", fertilizer, "Water", water, "Light", light, "Temperature", temperature, "Humidity", humidity, "Location", location)
		}
	}

	return closestLocation
}

func findClosestMin(value int, entityMap []EntityMapping) int {
	delta := 0
	resultId := -1

	for _, entity := range entityMap {
		if entity.Source <= value && value < entity.Source+entity.Range {
			delta = entity.Delta
		}
	}
	resultId = value + delta

	return resultId
}

// [
// 	{ "src": 0, "dest": -1, "delta": 0 },
// 	{ "src": 50, "dest": 52, "delta": 2 },
// 	{ "src": 98, "dest": 50, "delta": -48 },
//   { "src": 100, "delta": 0 }
// ]
// [
// 	{ "src": 0, "dest": 39, "delta": 39 },
// 	{ "src": 15, "delta": -15 },
// 	{ "src": 52, "delta": -15 },
// 	{ "src": 54, "delta": 0 }
// ]
// Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
// Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
// Seed 55, soil 57, fertilizer 57, water 53, light 46, temperature 82, humidity 82, location 86.
// Seed 13, soil 13, fertilizer 52, water 41, light 34, temperature 34, humidity 35, location 35.
