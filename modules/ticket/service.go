package ticket

import (
	"Parking_Slot_in_GO/modules/parking_spot"
	"Parking_Slot_in_GO/modules/vehicle"
	"Parking_Slot_in_GO/pkg/utitliy"
	"github.com/gin-gonic/gin"
	"log"
)

func BookSpotAndReturnTicket(ctx *gin.Context, vehicleObj *vehicle.Vehicle) {
	if vehicleObj == nil {
		utitliy.ClientErrResponse(ctx, "Invalid Request Body")
		return
	}

	var ticket *Ticket
	var parkingSpot *parking_spot.ParkingSpot
	parkingSpot, err := parking_spot.BookAndReturnParkingSpot(vehicleObj)
	if err != nil {
		utitliy.ClientErrResponse(ctx, "Invalid Request Body")
		return
	}

	if parkingSpot != nil {
		vehicle.SaveVehicleDetails(vehicleObj)

		ticket.ParkingSpotDetails = utitliy.GetJsonString(parkingSpot)
		ticket.VehicleDetails = utitliy.GetJsonString(vehicleObj)
		ticket.RatePerHrs = vehicleObj.RatePerHour
		ticket.BookingStartedAt = utitliy.GetCurTimeStamp()
		ticket.CreatedAt = utitliy.GetCurTimeStamp()
		ticket.Id = utitliy.GetUUID()

		utitliy.SuccessResponse(ctx, ticket, "Successfully Book a Slot. Here is your ticket. Parking spot number is =")
	} else {
		log.Printf("parking spot is nil")
	}

	utitliy.ServerErrResponse(ctx, "No Parking Spot Available Right Now.")
}
