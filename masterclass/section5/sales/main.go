package main

import (
	"fmt"
)

/*
Below is the code from the course. I do not like that the strings
are used to index the itemPrices map instead of an enum. I also don't
like that the sale determination is based on a string suffix

I have rewritten it below.
*/

/*
var itemPrices = map[string]float64{
	"TSHIRT": 20.00,
	"MUG":    12.50,
	"HAT":    18.00,
	"BOOK":   25.99,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := itemPrices[itemCode]
	if !found {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found = itemPrices[originalItemCode]
			if found {
				salePrice := basePrice * 0.9
				fmt.Printf(
					" - Item %s (Sale! Original: $%.2f, Sale Price: $%.2f)\n",
					originalItemCode, basePrice, salePrice,
				)
				return salePrice, true
			}
		}
		fmt.Printf(" - Item: %s (Product not found)\n", itemCode)
		return 0.0, false
	}
	fmt.Printf(" - Item: %s (Price: $%.2f)\n", itemCode, basePrice)
	return basePrice, found
}

func main() {
	orderItems := []string{
		"TSHIRT", "MUG_SALE", "HAT", "BOOK", "JAR",
	}
	var subtotal float64

	fmt.Println("------ Processing Order Items ------")

	for _, item := range orderItems {
		price, found := calculateItemPrice(item)
		if found {
			subtotal += price
		}
	}

	fmt.Printf("Subtotal: $%.2f\n", subtotal)
}
*/

//go:generate go tool stringer -type=Product

type Product int

const (
	Tshirt Product = iota
	Mug
	Hat
	Book
	Jar
)

var productPrices = map[Product]float64{
	Tshirt: 20.00,
	Mug:    12.50,
	Hat:    18.00,
	Book:   25.99,
}

func calculatePrice(product Product, sale bool) (float64, bool) {
	price, found := productPrices[product]
	if found && sale {
		salePrice := price * 0.9
		fmt.Printf(
			" - Item %v (Sale! Original: $%.2f, Sale Price: $%.2f)\n",
			product, price, salePrice,
		)
		return salePrice, found
	}
	if !found {
		fmt.Printf(" - Item: %v (Product not found)\n", product)
		return 0.0, found
	}
	fmt.Printf(" - Item: %v (Price: $%.2f)\n", product, price)
	return price, found
}

func main() {
	products := []Product{
		Tshirt, Mug, Hat, Book, Jar,
	}
	var subtotal float64

	fmt.Println("------ Processing Order Items ------")

	for _, product := range products {
		isOnSale := product == Mug
		price, found := calculatePrice(product, isOnSale)
		if found {
			subtotal += price
		}
	}

	fmt.Printf("Subtotal: $%.2f\n", subtotal)
}
