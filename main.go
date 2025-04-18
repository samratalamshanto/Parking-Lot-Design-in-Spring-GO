package main

import (
	"Parking_Slot_in_GO/pkg/config"
	"Parking_Slot_in_GO/pkg/database"
	"Parking_Slot_in_GO/pkg/global_exception_handler"
	"Parking_Slot_in_GO/pkg/redis"
	"Parking_Slot_in_GO/pkg/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	config.LoadConfig()
	database.ConnectDb()
	redis.ConnectRedis()

	routerEngine := gin.Default()

	routerEngine.Use(global_exception_handler.GlobalExceptionHandler())
	router.ConfigRouter(routerEngine)

	var PORT = os.Getenv("PORT")
	err := routerEngine.Run(":" + PORT)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	fmt.Println("Application running...")
}
