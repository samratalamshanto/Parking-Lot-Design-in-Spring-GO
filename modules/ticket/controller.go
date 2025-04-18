package ticket

import (
	"Parking_Slot_in_GO/modules/vehicle"
	"Parking_Slot_in_GO/pkg/utitliy"
	"github.com/gin-gonic/gin"
	"log"
)

func CheckoutTicket(ctx *gin.Context) {

}

func GetTicket(ctx *gin.Context) {
	var vehicleObj *vehicle.Vehicle
	err := ctx.ShouldBind(&vehicleObj)
	if err != nil {
		utitliy.ClientErrResponse(ctx, "Invalid Request Body.")
		log.Println("Bind error: %s", err.Error())
		ctx.Abort()
	}
	BookSpotAndReturnTicket(ctx, vehicleObj)
}
