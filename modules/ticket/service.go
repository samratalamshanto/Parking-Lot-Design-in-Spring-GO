package ticket

import (
	"Parking_Slot_in_GO/modules/parking_spot"
	"Parking_Slot_in_GO/modules/vehicle"
	"Parking_Slot_in_GO/pkg/database"
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
		ticket.RatePerHrs = vehicleObj.RatePerHrs
		ticket.BookingStartedAt = utitliy.GetCurTimeStamp()
		ticket.CreatedAt = utitliy.GetCurTimeStamp()
		ticket.Id = utitliy.GetUUID()

		err := saveTicket(ticket)
		if err != nil {
			log.Println(err)
			utitliy.ServerErrResponse(ctx, "Save Ticket Failed")
			return
		}

		utitliy.SuccessResponse(ctx, ticket, "Successfully Book a Slot. Here is your ticket. Parking spot number is =")
	}

	log.Printf("parking spot is nil")
	utitliy.SuccessResponse(ctx, nil, "No Parking Spot Available Right Now.")
}

func saveTicket(ticket *Ticket) error {
	query := `insert into ticket_details (id,created_at,booking_started_at,rate_per_hrs,vehicle_details,parking_spot_details) values (?, ?, ?, ?, ?, ?)`
	return database.DbSession.Query(query, ticket.Id, ticket.CreatedAt, ticket.BookingStartedAt, ticket.RatePerHrs, ticket.VehicleDetails, ticket.ParkingSpotDetails).Exec()
}
