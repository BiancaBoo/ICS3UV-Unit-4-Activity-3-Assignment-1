/**
 * @author Bianca Boo
 * @version 1.0.0
 * @date 2025-12-07
 * @fileoverview This program uses arrays to calculate totals, tax and discount of shopping cart. 
 */

// Constants
const TAX_RATE = 0.13;
const DISCOUNT_RATE = 0.10;
const DISCOUNT_LIMIT = 350.0;

// Arrays
const itemNames: string[] = [];
const itemPrices: number[] = [];

// Variables
let numItems: number = 0;
let subtotal: number = 0;
let discount: number = 0;
let tax: number = 0;
let total: number = 0;
let i: number = 0;

// Input: number of items
numItems = Number(prompt("How many items are you purchasing?") || "0");

// Input loop
for (i = 0; i < numItems; i++) {
  itemNames[i] = prompt("Enter the name of item:") || "";
  itemPrices[i] = Number(prompt("Enter the price of the item:") || "0");

  subtotal = subtotal + itemPrices[i];
}

// Discount check
if (subtotal > DISCOUNT_LIMIT) {
  discount = subtotal * DISCOUNT_RATE;
}

// Tax calculation
tax = (subtotal - discount) * TAX_RATE;

// Final total
total = (subtotal - discount) + tax;

// Output
console.log("Your shopping cart includes:");
console.log(itemNames.join(", "));

console.log("The total amount of items in your shopping cart is " + numItems);

console.log("The subtotal cost of your shopping trip was $" + subtotal.toFixed(2));

if (discount > 0) {
  console.log("You are eligible for a discount of $" + discount.toFixed(2));
} else {
  console.log("You are not eligible for a discount.");
}

console.log("The HST is $" + tax.toFixed(2));
console.log("The total is $" + total.toFixed(2));

console.log("Done.");