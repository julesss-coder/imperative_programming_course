// Problem description (in German)

// Aufgabe A3: Uhrzeiger
// Schreiben Sie ein Programm, das eine Uhrzeit (Stunde 0 bis 23, Minute 0 bis 59) einliest und den Winkel
// des Stunden- und Minutenzeigers zu dieser Uhrzeit ausgibt. Die Winkel sind 0°, wenn die Zeiger oben
// stehen, also um 0 und 12 Uhr.
// Die Ausgabe Ihres Programmes könnte so aussehen:
// Geben Sie bitte eine Uhrzeit an, indem Sie zunächst
// die Stunde (von 0 bis 23) und dann die Minute (von 0 bis 59) eingeben:
// 8
// 5
// Zeigerstellung um 08:05 Uhr:
// Winkel des Stundenzeigers: 242.5°
// Winkel des Minutenzeigers: 30°

//  Strategy 1

// Get hour
// Get minutes
// hour hand angle: starting position at given hour + movement during given minutes
// 30 deg * hour + (30 deg / 60 min) * minutes
// hour_hand_angle = 30 * (minutes/60 + hour)
// minute_hand_angle = 6 * minutes
// print hour_hand_angle
// print minute_hand_angle

// =================
// TODOS
//
// Throw error if hour and minutes input out of range
// Add error handling if input is not integer (even though data type is float).
// Consider edge cases.

package main

import (
	"fmt"
)

func get_clock_hand_angles_from_top(hour float64, minutes float64) {
	hour_hand_angle := 30 * (minutes/60 + hour)
	minute_hand_angle := 6 * minutes
	fmt.Printf("The angle of the hour hand from the 12 o'clock position is %.2f degrees. The angle of the minute hand from the same position is %.2f degrees.", hour_hand_angle, minute_hand_angle)
}

func main() {
	var hour float64
	var minutes float64

	fmt.Println("Enter a time for which to calculate the clock hand angles.")
	fmt.Println("Enter the hour (integer).")

	if _, error := fmt.Scan(&hour); error != nil {
		fmt.Println("Error occurred during input of hour.")
		return
	}

	if hour > 24 || hour < 0 {
		fmt.Println("The hours you entered are invalid.")
		return
	}

	fmt.Println("Enter the minutes (integer).")

	if _, error := fmt.Scan(&minutes); error != nil {
		fmt.Println("Error occurred during input of minute.")
		return
	}

	if minutes < 0 || minutes > 59 {
		fmt.Println("The minutes you entered are invalid.")
		return
	}

	get_clock_hand_angles_from_top(hour, minutes)
}
