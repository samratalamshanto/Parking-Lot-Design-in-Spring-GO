package router

import (
	"Parking_Slot_in_GO/modules/parking_spot"
	"Parking_Slot_in_GO/modules/ticket"
	"github.com/gin-gonic/gin"
)

func ConfigRouter(ginRouter *gin.Engine) {
	parking_spot.SetupRouter(ginRouter)
	ticket.SetupRouter(ginRouter)
}
