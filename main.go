package main

import (
	"Parking_Slot_in_GO/pkg/config"
	"Parking_Slot_in_GO/pkg/database"
	"Parking_Slot_in_GO/pkg/redis"
	"fmt"
)

func main() {
	config.LoadConfig()
	database.ConnectDb()
	redis.ConnectRedis()
	fmt.Println("Application running...")
}
