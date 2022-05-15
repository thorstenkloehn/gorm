package controller

import (
	"Gorm_uebung/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Controller struct {
	model.Datenbank
}

func (c *Controller) Index() *[]model.Datenbank {
	datenbank, _ := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
	var datenbank1 []model.Datenbank
	datenbank.Find(&datenbank1)
	return &datenbank1
}

func (c *Controller) Create() {
	datenbank, _ := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
	var datenbank1 = model.Datenbank{Titel: c.Titel, Name: c.Name}
	datenbank.Create(&datenbank1)
}

func (c *Controller) Store() {

}
