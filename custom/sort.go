package custom

import (
	"math"
	"sort"
)

func Sort1(results []map[int]int) {
	// Sort based on the lowest key having the highest value
	sort.Slice(results, func(i, j int) bool {
		// For comparison, get the values of the lowest key (1 in this case)
		// Adjust the key based on your use case if needed
		for key := range results[i] {
			if valueI, okI := results[i][key]; okI {
				if valueJ, okJ := results[j][key]; okJ {
					return valueI > valueJ
				}
			}
		}
		return false
	})
}

func Sort2(results []map[int]int) {
	// Function to calculate the range (max-min) of values in a map
	calculateRange := func(m map[int]int) int {
		minVal, maxVal := math.MaxInt32, math.MinInt32
		for _, v := range m {
			if v < minVal {
				minVal = v
			}
			if v > maxVal {
				maxVal = v
			}
		}
		return maxVal - minVal
	}

	// Sort the results based on the range of values in each map (min-max difference)
	sort.Slice(results, func(i, j int) bool {
		return calculateRange(results[i]) < calculateRange(results[j])
	})
}

// Function to calculate how well the map follows the "lower keys have higher values" rule
func lowerKeysHigherValuesScore(m map[int]int) int {
	score := 0
	for key1, val1 := range m {
		for key2, val2 := range m {
			// We want lower keys (key1 < key2) to have higher values (val1 > val2)
			if key1 < key2 && val1 > val2 {
				score += val1 - val2
			}
		}
	}
	return score
}

// Function to calculate the range (max-min) of values in a map
func calculateRange(m map[int]int) int {
	minVal, maxVal := math.MaxInt32, math.MinInt32
	for _, v := range m {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal - minVal
}

func Sort3(results []map[int]int) {
	// Sort the results based on combined balance and lower-keys-higher-values rule
	sort.Slice(results, func(i, j int) bool {
		// Calculate the score for how well the map follows the "lower keys higher values" rule
		scoreI := lowerKeysHigherValuesScore(results[i])
		scoreJ := lowerKeysHigherValuesScore(results[j])

		// Calculate the balance (range of values) for each map
		rangeI := calculateRange(results[i])
		rangeJ := calculateRange(results[j])

		// Combine both scores:
		// - prioritize lower-keys-higher-values (higher scoreI)
		// - also prefer a smaller range for more balanced distribution (lower rangeI)
		return (scoreI - rangeI) > (scoreJ - rangeJ)
	})
}
