package main

import (
	"errors"
	"fmt"

	"github.com/piyushsingariya/pokerchange/custom"
	"github.com/sirupsen/logrus"
)

func Ptr[T any](v T) *T {
	return &v
}

type Chip struct {
	Color string `json:"color"`
	Value int    `json:"value"`
}

type Distribution struct {
	Chip
	Number int `json:"number"`
}

var (
	chips = []*Chip{
		{
			Color: "Black",
			Value: 100,
		},
		{
			Color: "Green",
			Value: 50,
		},
		{
			Color: "Blue",
			Value: 25,
		},
		{
			Color: "Red",
			Value: 10,
		},
		{
			Color: "White",
			Value: 5,
		},
	} // Available coin denominations
)

func combinations(chips []*Chip, amount int, maxCount int) ([][]Distribution, error) {
	// Map to keep track of coin counts
	coinCounts := make(map[int]int)
	var results []map[int]int

	coins := []int{}
	// Calculate the total number of coins required to meet the amount
	for _, chip := range chips {
		coin := chip.Value
		// Use each coin at least once
		coinCounts[coin] = 1
		coins = append(coins, coin)
		amount -= coin
	}

	// If amount is negative, it's impossible to form the amount
	if amount < 0 {
		return nil, errors.New("amount can not be distributed")
	}

	// Start the backtracking process
	backtrack(coins, amount, maxCount, 0, coinCounts, &results)

	custom.Sort3(results)

	final := [][]Distribution{}
	for _, result := range results {
		converted := []Distribution{}
		for _, chip := range chips {
			converted = append(converted, Distribution{
				Chip:   *chip,
				Number: result[chip.Value],
			})
		}

		final = append(final, converted)
	}
	return final, nil
}

// backtrack function to find a combination of coins
func backtrack(coins []int, amount int, maxCount int, index int, coinCounts map[int]int, result *[]map[int]int) {
	// Check if we have reached the target amount
	if amount == 0 {
		// If the amount is exactly zero, store the current counts of coins
		*result = append(*result, copyCounts(coinCounts))
		return
	}
	// If the amount is negative or we have exhausted all coins
	if amount < 0 || index >= len(coins) {
		return
	}

	// Try using the current coin
	coin := coins[index]
	if amount >= coin && coinCounts[coin] < maxCount {
		coinCounts[coin]++
		// Recur with the reduced amount and the same index to reuse the same coin
		backtrack(coins, amount-coin, maxCount, index, coinCounts, result)
		coinCounts[coin]--
	}

	// Try skipping the current coin and moving to the next one
	backtrack(coins, amount, maxCount, index+1, coinCounts, result)
}

// Helper function to copy the counts to avoid reference issues
func copyCounts(original map[int]int) map[int]int {
	copy := make(map[int]int)
	for k, v := range original {
		copy[k] = v
	}
	return copy
}

func generateCoins(players int, amount int, maxChipPerColor int) {
	totalBuyIn := players * amount
	for i := 1; ; i = i * 2 {
		totalBankAmount := 0
		for _, chip := range chips {
			chip.Value = chip.Value * i
			totalBankAmount += chip.Value * maxChipPerColor
		}

		// fmt.Println(totalBankAmount)

		if totalBankAmount > int(1.5*float64(totalBuyIn)) {
			// json.NewEncoder(os.Stdout).Encode(chips)
			// fmt.Println(i, totalBankAmount, totalBuyIn)
			return
		}
	}
}

func main() {
	amount := 600  // Total amount to achieve
	maxCount := 60 // Maximum count for each coin
	players := 8
	generateCoins(players, amount, maxCount)
	result, err := combinations(chips, amount, maxCount/players)
	if err != nil {
		logrus.Fatal(err)
	}

	if len(result) > 0 {
		fmt.Println("")
		fmt.Println("Number of players in the game: ", players)
		fmt.Println("Buy-in for individual players: ", amount)
		fmt.Println("Possible combinations of coins and their counts:")
		for _, valued := range result[0] {
			fmt.Printf("Color %s, Coin: %d, Count: %d\n", valued.Color, valued.Value, valued.Number)
		}
		fmt.Println("----")
		// for _, counts := range result {
		// 	for _, valued := range counts {
		// 		fmt.Printf("Color %s, Coin: %d, Count: %d\n", valued.Color, valued.Value, valued.Number)
		// 	}
		// 	fmt.Println("----")
		// }
	} else {
		fmt.Println("It's not possible to make the given amount with the provided coins.")
	}

	// mid := len(result) / 2
	// if len(result)<<1 == 0 {
	// 	mid = mid + 1
	// }
	// json.NewEncoder(os.Stdout).Encode(result[mid])
	// json.NewEncoder(os.Stdout).Encode(result[mid-1])
	// json.NewEncoder(os.Stdout).Encode(result[mid+1])
	// json.NewEncoder(os.Stdout).Encode(result[1])
	// json.NewEncoder(os.Stdout).Encode(result[2])
}
