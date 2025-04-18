package vehicle

import (
	"Parking_Slot_in_GO/base/enums/colors"
	"Parking_Slot_in_GO/base/enums/vehicle"
	"Parking_Slot_in_GO/base/models"
)

type Vehicle struct {
	models.CommonFields
	LicencePlateNumber   string `json:"licencePlateNumber" gocql:"column:licence_plate_number"`
	vehicle.EVehicleType `json:"vehicleType" gocql:"column:vehicle_type"`
	colors.EColor        `json:"color" gocql:"column:color"`
	RatePerHrs           float64 `json:"ratePerHour" gocql:"column:rate_per_hrs"`
	TotalFee             float64 `json:"totalFee" gocql:"column:total_fee"`
	TotalHr              float64 `json:"totalHr" gocql:"column:total_hr"`
}
