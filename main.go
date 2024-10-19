// Problem description (in German):

// Aufgabe A1: Längenmaße
// Schreiben Sie ein Programm, das eine Länge in Metern einliest und diese umrechnet in:
// • mm
// • km
// • Zoll (1 Zoll = 2.54 cm)
// • Seemeilen (1 sm = 1852 m)
// • Lichtjahre (1 Lj = 9 460 730 472 580 800 m)
// Die Ausgabe Ihres Programmes könnte so aussehen:
// Geben Sie eine Länge (in Metern) ein:
// 3755
// 3755 Meter entsprechen:
// 3.755e+06 mm
// 3.755 km
// 147834.646 Zoll
// 2.03 sm
// 3.97e-13 Lj

package main

import (
	"fmt"
)

func convert_meters(meters float64) {
	var millimeters = meters * 1000
	fmt.Printf("%.2f meters are %f millimeters.\n", meters, millimeters)

	var kilometers = meters / 1000
	fmt.Printf("%.2f meters are %.3f kilometers.\n", meters, kilometers)

	var inches float64 = meters / 0.0254
	fmt.Printf("%.2f meters are %.2f inches.\n", meters, inches)

	var nautical_miles float64 = meters / 1852
	fmt.Printf("%.2f meters are %.6g nautical miles.\n", meters, nautical_miles)

	var light_years = meters / 9.461e15
	fmt.Printf("%.2f meters are %.15g light years.\n", meters, light_years)
}

func main() {
	var num float64
	var error error

	fmt.Println("Enter the number of meters you wish to convert to various formats.")
	_, error = fmt.Scan(&num)

	if error != nil {
		fmt.Println("Error reading user input: ", error)
	} else {
		fmt.Printf("You successfully entered the number %f. Converting to various units now...\n", num)
		convert_meters(num)
	}
}
