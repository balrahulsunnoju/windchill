// Package windchill provides functionality related to calculating wind chill factors.
package main

import (
	"flag"
	"fmt"
	"math"
)

func windChill(ta float64, v float64) float64 {
	return 35.74 + 0.6215*ta - 35.75*math.Pow(v, 0.16) + 0.4275*ta*math.Pow(v, 0.16)
}

func promptForInput() (float64, float64) {
	var ta, v float64
	for {
		fmt.Print("Enter the temperature (°F) between -58 and 41: ")
		_, err := fmt.Scan(&ta)
		if err != nil || ta < -58 || ta > 41 {
			fmt.Println("Invalid input. Please enter a valid temperature.")
			continue
		}

		fmt.Print("Enter the wind speed (mph) greater than or equal to 2: ")
		_, err = fmt.Scan(&v)
		if err != nil || v < 2 {
			fmt.Println("Invalid input. Please enter a valid wind speed.")
			continue
		}

		break
	}
	return ta, v
}

func printTable() {
	// Print table header
	fmt.Printf("%-4s", "V")
	for ta := 40.0; ta >= -45.0; ta -= 5 { // Changed the order for vertical alignment
		fmt.Printf("%8.1f", ta)
	}
	fmt.Println()

	// Print table rows
	for v := 5.0; v <= 60.0; v += 5 {
		fmt.Printf("%-4.0f", v)
		for ta := 40.0; ta >= -45.0; ta -= 5 {
			wc := windChill(ta, v)
			fmt.Printf("%8.1f", wc) // Display the wind chill value aligned to tenths place
		}
		fmt.Println()
	}
}

func main() {
	tableFlag := flag.Bool("table", false, "Display the wind chill table")
	flag.Parse()

	if *tableFlag {
		printTable()
		return
	}

	ta, v := promptForInput()
	wc := windChill(ta, v)
	fmt.Printf("The wind chill temperature is: %.2f°F\n", wc)
}
