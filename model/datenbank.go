package model

type Datenbank struct {
	Id    uint64 `gorm:"primaryKey"`
	Titel string
	Name  string
}
