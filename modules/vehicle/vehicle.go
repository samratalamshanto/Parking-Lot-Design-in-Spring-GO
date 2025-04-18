package vehicle

import (
	"Parking_Slot_in_GO/base/enums/colors"
	"Parking_Slot_in_GO/base/enums/vehicle"
	"Parking_Slot_in_GO/base/models"
)

type Vehicle struct {
	models.CommonFields
	LicencePlateNumber   string `json:"licencePlateNumber" gorm:"column:licencePlateNumber"`
	vehicle.EVehicleType `json:"vehicleType" gorm:"column:vehicleType"`
	colors.EColor        `json:"color" gorm:"column:color"`
	RatePerHour          float64 `json:"ratePerHour" gorm:"column:ratePerHour"`
	TotalFee             float64 `json:"totalFee" gorm:"column:totalFee"`
	TotalHr              float64 `json:"totalHr" gorm:"column:totalHr"`
}
