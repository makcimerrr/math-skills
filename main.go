package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run your-program.go data.txt")
		return
	}

	filePath := os.Args[1]

	// Lire le fichier
	data, err := readDataFromFile(filePath)
	// Si le fichier data n'exitse pas
	if err != nil {
		log.Fatalf("Error reading data from file: %v", err)
	}
	// Si data est vide
	if len(data) == 0 {
		log.Fatal("Error: The data file is empty.")
	}

	// Calculate average
	average := calculateAverage(data)
	fmt.Println("Average:", round(average)) // Appel round pour arrondir au supérieur

	// Calculate median
	median := calculateMedian(data)
	fmt.Println("Median:", round(median)) // Appel round pour arrondir au supérieur

	// Calculate variance
	variance := calculateVariance(data)
	fmt.Println("Variance:", round(variance)) // Appel round pour arrondir au supérieur

	// Calculate standard deviation
	standardDeviation := calculateStandardDeviation(data)
	fmt.Println("Standard Deviation:", round(standardDeviation)) // Appel round pour arrondir au supérieur
}

// Fonction pour lire les données du fichier
func readDataFromFile(filePath string) ([]float64, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var data []float64

	for _, line := range lines {
		if line != "" {
			value, err := strconv.ParseFloat(line, 64)
			if err != nil {
				return nil, err
			}
			data = append(data, value)
		}
	}

	return data, nil
}

// Function to calculate the average (moyenne)
func calculateAverage(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// Function to calculate the median (médiane)
func calculateMedian(data []float64) float64 {
	sort.Float64s(data)
	length := len(data)
	if length%2 == 0 {
		return (data[length/2-1] + data[length/2]) / 2
	}
	return data[length/2]
}

// Function to calculate the variance (variance)
func calculateVariance(data []float64) float64 {
	average := calculateAverage(data)
	variance := 0.0
	for _, value := range data {
		variance += math.Pow(value-average, 2)
	}
	return variance / float64(len(data))
}

// Function to calculate the standard deviation (écart type)
func calculateStandardDeviation(data []float64) float64 {
	variance := calculateVariance(data)
	return math.Sqrt(variance)
}

// Fonction pour arrondir la valeur float64 à l'entier le plus proche
func round(value float64) int {
	return int(math.Round(value))
}
