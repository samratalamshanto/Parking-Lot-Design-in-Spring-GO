package main

import (
	"Parking_Slot_in_GO/pkg/config"
	"Parking_Slot_in_GO/pkg/database"
	"fmt"
)

func main() {
	config.LoadConfig()
	database.ConnectDb()
	fmt.Println("Application running...")
}
