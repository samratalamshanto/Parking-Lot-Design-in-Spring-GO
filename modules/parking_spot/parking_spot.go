package parking_spot

import (
	"Parking_Slot_in_GO/modules/base/enums/parking_spot"
	"Parking_Slot_in_GO/modules/base/models"
)

type ParkingSpot struct {
	models.CommonFields
	SpotNumber                    string `json:"spot_number" gocql:"spot_number"`
	FloorNumber                   int    `json:"floor_number" gocql:"floor_number"`
	IsBooked                      bool   `json:"is_booked" gocql:"is_booked"`
	parking_spot.EParkingSpotType `json:"parking_spot_type" gocql:"parking_spot_type"`
}
