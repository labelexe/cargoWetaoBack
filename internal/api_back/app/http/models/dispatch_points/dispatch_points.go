package dispatch_points

import "Wetao/internal/database"

// DispatchPoint represents the dispatch_points table in the database
type DispatchPoint struct {
	ID          uint   `json:"-" gorm:"primarykey"`
	Name        string `json:"name"`
	Loc         string `json:"loc"`
	Lon         string `json:"lon"`
	FullAddress string `json:"full_address"`

	// Additional Gorm tags can be added if needed
}

// TableName sets the table name for the DispatchPoint model
func (DispatchPoint) TableName() string {
	return "dispatch_points"
}

func CreateNewDispatchPoints() {

}

func GetDispatchPoints() ([]DispatchPoint, error) {
	db := database.GetDB()
	var dispatchPoints []DispatchPoint
	//
	err := db.
		Select("*").
		Find(&dispatchPoints).Error
	return dispatchPoints, err
}
