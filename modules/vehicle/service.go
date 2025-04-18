package vehicle

import (
	"Parking_Slot_in_GO/pkg/database"
	"Parking_Slot_in_GO/pkg/utitliy"
	"log"
)

func SaveVehicleDetails(vehicle *Vehicle) {
	vehicle.Id = utitliy.GetUUID()
	err := saveVehicle(vehicle)
	if err != nil {
		log.Printf("save vehicle failed: %v\n", err)
	}
}

func saveVehicle(vehicle *Vehicle) error {
	query := `insert into vehicle_details (id,created_at,rate_per_hrs, licence_plate_number,vehicle_type, color) values (?, ?, ?, ?, ?, ?)`
	return database.DbSession.Query(query, vehicle.Id, utitliy.GetCurTimeStamp(), vehicle.RatePerHrs, vehicle.LicencePlateNumber,
		vehicle.EVehicleType, vehicle.EColor).Exec()
}
