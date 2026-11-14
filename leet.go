package main 

//Naive approach
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

//Optimized approach
func maxProfitV2(prices []int) int {
    maxProfit := 0
    minPrice := prices[0]

    for i := 0; i < len(prices); i++ {
        if prices[i] - minPrice > maxProfit {
            maxProfit = prices[i] - minPrice
        }
        if prices[i] < minPrice {
            minPrice = prices[i] 
        }
    }

    return maxProfit
}