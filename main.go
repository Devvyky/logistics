package main

import (
	"fmt"
	"sort"
)

// Imagine for a moment that one of our product lines ships in various pack sizes:
//  250 Items
//  500 Items
//  1000 Items
//  2000 Items
//  5000 Items

// Our customers can order any number of these items through our website, but they will
// always only be given complete packs.

// 1. Only whole packs can be sent. Packs cannot be broken open.
// 2. Within the constraints of Rule 1 above, send out no more items than necessary to
// fulfil the order.
// 3. Within the constraints of Rules 1 &amp; 2 above, send out as few packs as possible to
// fulfil each order.

// So, for example:
// Items ordered: 1
// Correct number of packs: 1 x 250
// Incorrect number of packs: 1 x 500 – more items than necessary

// Items ordered: 250
// Correct number of packs: 1 x 250
// Incorrect number of packs: 1 x 500 – more items than necessary

// Items ordered: 251
// Correct number of packs: 1 x 500
// Incorrect number of packs: 2 x 250 – more packs than necessary

// Items ordered: 501
// Correct number of packs: 1 x 500, 1 x 250
// Incorrect number of packs: 1 x 1000 – more items than necessary

// Items ordered: 12001
// Correct number of packs: 2 x 5000, 1 x 2000, 1 x 250
// Incorrect number of packs: 3 x 5000 – more items than necessary

// Write a program that will calculate the number of packs of each pack size that should be sent
func main() {
	fmt.Println(calculatePacks(1))
}

func calculatePacks(items int) map[int]int {
	packs := map[int]int{
		250:  0,
		500:  0,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	// sort the packs in descending order
	packSizes := []int{5000, 2000, 1000, 500, 250}
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	// for each pack size, calculate the number of packs required
	for _, packSize := range packSizes {
		// if the number of items is less than the pack size, continue to the next pack size
		if items < packSize {
			continue
		}

		// calculate the number of packs required
		numPacks := items / packSize

		// update the packs map with the number of packs required
		packs[packSize] = numPacks

		// update the remaining items
		items -= numPacks * packSize
	}

	// return the packs map
	return packs
}
