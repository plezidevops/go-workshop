package room

import "fmt"

const roomText = "%d : %d people / %d nights"

// display information about a room
func PrintDetails(roomNumber, size, nights int) {
	fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
}
