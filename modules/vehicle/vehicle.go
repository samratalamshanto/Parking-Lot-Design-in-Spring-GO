package vehicle

import (
	"Parking_Slot_in_GO/modules/base/enums"
	"Parking_Slot_in_GO/modules/base/enums/colors"
	"Parking_Slot_in_GO/modules/base/models"
)

type Vehicle struct {
	models.CommonFields
	LicencePlateNumber string `json:"licencePlateNumber" gorm:"column:licencePlateNumber"`
	enums.EVehicleType `json:"vehicleType" gorm:"column:vehicleType"`
	colors.EColor      `json:"color" gorm:"column:color"`
	RatePerHour        float32 `json:"ratePerHour" gorm:"column:ratePerHour"`
	TotalFee           float32 `json:"totalFee" gorm:"column:totalFee"`
	TotalHr            float32 `json:"totalHr" gorm:"column:totalHr"`
}
