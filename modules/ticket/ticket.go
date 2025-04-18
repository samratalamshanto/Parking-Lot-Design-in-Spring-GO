package ticket

import (
	"Parking_Slot_in_GO/modules/base/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Ticket struct {
	models.CommonFields
	BookingStartedAt   timestamppb.Timestamp `json:"booking_started_at" gocql:"booking_started_at"`
	BookingEndedAt     timestamppb.Timestamp `json:"booking_ended_at" gocql:"booking_ended_at"`
	TotalHrs           float64               `json:"total_hrs" gocql:"total_hrs"`
	TotalFees          float64               `json:"total_fees" gocql:"total_fees"`
	RatePerHrs         float64               `json:"rate_per_hrs" gocql:"rate_per_hrs"`
	VehicleDetails     string                `json:"vehicle_details" gocql:"vehicle_details"`
	ParkingSpotDetails string                `json:"parking_spot_details" gocql:"parking_spot_details"`
}
