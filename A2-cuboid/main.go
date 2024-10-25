// Problem description (in German)

// Aufgabe A2: Quader
// Schreiben Sie ein Programm, daß die Längen der drei Seiten eines Quaders einliest und folgende
// Größen dieses Quaders ausgibt:
// • Volumen
// • Kantensumme
// • Oberfläche
// • Umkugelradius
// • Raumdiagonale
// Tipps: Wenn Ihnen die Formeln zur Bestimmung dieser Größen nicht bekannt sind, finden Sie sie in einer
// mathematischen Formelsammlung oder auch der Wikipedia. Verwenden Sie geeignete Funktionen
// aus dem Package math.
// Die Ausgabe des Programmes könnte so aussehen:
// Bitte geben Sie die drei Seitenlängen des Quaders ein:
// 3
// 4.5
// 8
// Ein 3 × 4.5 × 8 Quader hat die geometrischen Größen:
// Volumen: 108
// Kantensumme: 62
// Oberfläche: 147
// Umkugelradius: 4.83
// Raumdiagonale: 9.66

package main

import (
	"fmt"
	"math"
)

func calculate_cuboid(a float64, b float64, c float64) {
	var volume float64 = a * b * c
	var sum_of_sides float64 = 4*a + 4*b + 4*c
	var surface float64 = 2*a*c + 2*b*c + 2*a*b
	var circumradius float64 = float64(1 / (2 * math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2)+math.Pow(c, 2))))
	var space_diagonal float64 = float64(math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2) + math.Pow(c, 2)))

	fmt.Printf(`For a cuboid with sidelengths a = %.2f, b = %.2f, and c = %.2f, the volume is %.2f m^3, the sum of its sides is %.2f m, its surface area is %.2f m^2, its circumradius is %.2f and its space diagonale is %.2f.`, a, b, c, volume, sum_of_sides, surface, circumradius, space_diagonal)
}

func main() {
	var side_a float64
	var side_b float64
	var side_c float64
	var error error

	fmt.Println("Enter a side length:")
	_, error = fmt.Scan(&side_a)
	if error != nil {
		fmt.Println("Error reading input.")
	}
	fmt.Println("Enter a side length:")
	_, error = fmt.Scan(&side_b)
	if error != nil {
		fmt.Println("Error reading input.")
	}
	fmt.Println("Enter a side length:")
	_, error = fmt.Scan(&side_c)
	if error != nil {
		fmt.Println("Error reading input.")
	}

	fmt.Printf("a, b, c: %.2f, %.2f, %.2f\n", side_a, side_b, side_c)
	calculate_cuboid(side_a, side_b, side_c)

}
