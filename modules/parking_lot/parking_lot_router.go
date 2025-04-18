package parking_lot

import "github.com/gin-gonic/gin"

func SetupRouter(ginRouter *gin.Engine) {
	parkingSlot := ginRouter.Group("parking-slot")
	{
		parkingSlot.POST("/add")
		parkingSlot.POST("/delete")
		parkingSlot.GET("/free-parking-slot")
		parkingSlot.POST("/book")
	}
}
