package main 

func maxProfit(prices []int) int {
    results := []int{}

    for i := 0; i < len(prices); i++ {
        for j := i + 1; j < len(prices); j++ {
            profit := prices[j] - prices[i]
            results = append(results, profit)
        }
    }

    maxProfit := 0                  
    for _, p := range results {
        if p > maxProfit {
            maxProfit = p
        }
    }

    return maxProfit
}