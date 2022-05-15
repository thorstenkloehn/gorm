package main

import (
	"Gorm_uebung/controller"
	"Gorm_uebung/model"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Datenbank{})
	// Test
	var xxx = &controller.Controller{}
	ausgabe := *xxx.Index()
	fmt.Println(ausgabe)

	var ttt controller.Controller
	ttt.Titel = "thorstenkkkkk"
	ttt.Name = "Pjkkljkklkjl"
	ttt.Create()
}
