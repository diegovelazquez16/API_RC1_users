package models

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Genero    string `json:"genero"` 
}

