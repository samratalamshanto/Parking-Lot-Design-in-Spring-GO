package router

import (
	"Parking_Slot_in_GO/modules/parking_lot"
	"Parking_Slot_in_GO/modules/ticket"
	"github.com/gin-gonic/gin"
)

func ConfigRouter(ginRouter *gin.Engine) {
	parking_lot.SetupRouter(ginRouter)
	ticket.SetupRouter(ginRouter)
}
