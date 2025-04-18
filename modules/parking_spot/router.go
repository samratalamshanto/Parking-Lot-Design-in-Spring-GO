package parking_spot

import "github.com/gin-gonic/gin"

func SetupRouter(ginRouter *gin.Engine) {
	parkingSlot := ginRouter.Group("parking-slot")
	{
		parkingSlot.POST("/add", AddParkingSpot)
		parkingSlot.POST("/delete", DeleteParkingSpot)
		parkingSlot.GET("/free-parking-slot", FreeParkingSpot)
		parkingSlot.POST("/book", BookParkingSpot)
	}
}
