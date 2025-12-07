// Author: Bianca Boo
// Version: 1.0.0
// Date: 2025-12-02
// Fileoverview: This program uses arrays to calculate totals, tax and discount of shopping cart.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const TAX_RATE float64 = 0.13
	const DISCOUNT_RATE float64 = 0.10
	const DISCOUNT_LIMIT float64 = 350.0

	var numItems int
	var subtotal float64
	var discount float64
	var tax float64
	var total float64

	scanner := bufio.NewScanner(os.Stdin)

	// Get number of items
	for {
		fmt.Print("How many items are you purchasing? ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		n, err := strconv.Atoi(input)
		if err != nil || n <= 0 {
			fmt.Println("Please enter a valid positive number.")
			continue
		}
		numItems = n
		break
	}

	// Arrays
	itemNames := make([]string, numItems)
	itemPrices := make([]float64, numItems)

	// Input loop
	for i := 0; i < numItems; i++ {
		fmt.Printf("Enter the name of item #%d: ", i+1)
		scanner.Scan()
		itemNames[i] = strings.TrimSpace(scanner.Text())

		for {
			fmt.Printf("Enter the price of item #%d: ", i+1)
			scanner.Scan()
			priceStr := strings.TrimSpace(scanner.Text())
			price, err := strconv.ParseFloat(priceStr, 64)
			if err != nil || price < 0 {
				fmt.Println("Please enter a valid non-negative number.")
				continue
			}
			itemPrices[i] = price
			break
		}

		subtotal += itemPrices[i]
	}

	// Discount check
	if subtotal > DISCOUNT_LIMIT {
		discount = subtotal * DISCOUNT_RATE
	}

	// Tax calculation
	tax = (subtotal - discount) * TAX_RATE
	total = (subtotal - discount) + tax

	// Output
	fmt.Print("\nYour shopping cart includes: ")
	for i := 0; i < numItems; i++ {
		if i < numItems-1 {
			fmt.Print(itemNames[i] + ", ")
		} else {
			fmt.Println(itemNames[i])
		}
	}

	fmt.Printf("The total amount of items in your shopping cart is %d\n", numItems)

	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	if discount > 0 {
		fmt.Printf("Discount: $%.2f\n", discount)
	} else {
		fmt.Println("No discount applied.")
	}
	fmt.Printf("HST: $%.2f\n", tax)
	fmt.Printf("Total: $%.2f\n", total)

	fmt.Println("\nDone.")
}
