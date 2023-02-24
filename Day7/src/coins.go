package coins

import (
	"sync"
)

// MinCoins is a function that calculates the minimum number of coins
// needed to make a given value.
// The function takes in two arguments:
//   val - the value to make
//   coins - a slice of coin values available
// It returns a slice of coin values that add up to the given value.
func MinCoins(val int, coins []int) []int {
    // Initialize an empty slice to store the result
    res := make([]int, 0)
    // Start from the highest denomination coin and work down
    i := len(coins) - 1
    for i >= 0 {
        // Keep subtracting the coin from the value as long as it's less than the value
        for val >= coins[i] {
            val -= coins[i]
            res = append(res, coins[i])
        }
        // Move on to the next lower denomination coin
        i -= 1
    }
    // Return the result
    return res
}


// MinCoins2 is a function that calculates the minimum number of coins required to make change for a given value.
func MinCoins2(val int, coins []int) []int {
    // minCoins is a slice that stores the minimum number of coins required to make change for each value from 0 to val.
    minCoins := make([]int, val+1)

    // Loop through each value from 1 to val.
    for i := 1; i <= val; i++ {
        // Initialize the number of coins required to make change for i as val + 1 (a value that is guaranteed to be higher than any possible number of coins).
        minCoins[i] = val + 1

        // Loop through each coin in the coins slice.
        for j := 0; j < len(coins); j++ {
            // If a coin has a value of 0, return an empty slice (as change cannot be made with a 0-valued coin).
            if coins[j] <= 0 {
                return []int{}
            }

            // If the current coin is less than or equal to the current value, check if using it would result in a lower number of coins than the current minimum.
            if coins[j] <= i {
                if minCoins[i-coins[j]]+1 < minCoins[i] {
                    // If using the current coin would result in a lower number of coins, update the minimum number of coins required to make change for the current value.
                    minCoins[i] = minCoins[i-coins[j]] + 1
                }
            }
        }
    }

    // change is a slice that will store the actual coins used to make change for the given value.
    change := []int{}

    // Loop through each value from val down to 0.
    for i := val; i > 0; {
        // Store the current value in a separate variable.
        t := i

        // Loop through each coin in the coins slice.
        for j := 0; j < len(coins); j++ {
            // If the current coin is less than or equal to the current value, and using it would result in the minimum number of coins, add the coin to the change slice and subtract its value from the current value.
            if coins[j] <= i && minCoins[i-coins[j]]+1 == minCoins[i] {
                change = append(change, coins[j])
                i -= coins[j]
                break
            }
        }

        // If the current value has not changed, break out of the loop.
        if t == i {
            break
        }
    }

    // Return the change slice.
    return change
}

var mu1 sync.Mutex
var mu2 sync.Mutex

// chacgeArr is a function to change the value of the slice pointed to by `t` to `val`.
// It acquires a lock on `mu1` before making the change to ensure mutual exclusion.
func chacgeArr(t *[]int, val []int) {
    mu1.Lock()
	defer mu1.Unlock()
    *t = val
}

// chacgeVal is a function to change the value of the integer pointed to by `t` to `val`.
// It acquires a lock on `mu2` before making the change to ensure mutual exclusion.
func chacgeVal(t *int, val int) {
    mu2.Lock()
	defer mu2.Unlock()
    *t = val
}

// MinCoins4 returns the minimum number of coins needed to reach the given price
func MinCoins4(price int, denoms []int) (r []int) {
    var wg sync.WaitGroup
    count := price + 1
    c := make([]int, count)
    // Loop through all the denominations
    for _,i := range denoms {
        if i < price {
            wg.Add(1)
            // For each denomination less than the price, call helper4
            go helper4(&count, denoms, price - i, &wg, &c, i)
        } else if i == price {
            // If a denomination equals the price, return it
            return []int{i}
        }
    }
    wg.Wait()
    // Return the result
    return c
}

// helper4 calls recurs4 and updates the result slice if the returned value is true
func helper4(count *int, denoms []int, price int, wg *sync.WaitGroup, c *[]int, i int) {
    defer wg.Done()
    var res []int
    if recurs4(count, denoms, price, 1, &res) {
        res = append(res, i)
        chacgeArr(c, res)
    }
}

// recurs4 is a recursive function that finds the minimum number of coins needed to reach the given price
func recurs4(count *int, denoms []int, price, depth int, res *[]int) bool {
    // If the depth is greater than the current count, return false
    if *count <= depth {
        return false
    }
    f := false
    var rr []int
    // Loop through all the denominations
    for _,i := range denoms {
        if i < price {
            rrr := []int{}
            // Recursively call recurs4 for the remaining price
            r := recurs4(count, denoms, price - i, depth + 1, &rrr)
            if r {
                rr = append(rrr, i)
                f = true
            }
        } else if i == price {
            *res = append(*res, i)
            // Update the count if a denomination equals the price
            chacgeVal(count, depth)
            return true
        }
        // If the depth is greater than the current count, return false
        if *count <= depth {
            return false
        }
    }
    if f {
        *res = rr
    }
    return f
}







