package cargo

func (Type) TableName() string {
	return "cargo_types"
}

// Type Типы (карго)
type Type struct {
	ID   uint   `json:"-" gorm:"primaryKey"`
	Name string `json:"name" gorm:"name"`
	Slug string `json:"slug"`
}
