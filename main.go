package main

import (
	"database/sql"
	"klik/controler"
	"klik/handler"
	"klik/routes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mnurfilza/lib"
)

var (
	user     = "root"
	password = "mypass123"  // adjust with your mysql configuration
	host     = "172.17.0.2" //adjust with your configuration
	port     = "3306"
	dbname   = "klik_test"
)

func main() {
	db, err := InitDatabase()
	if err != nil {
		db, err = lib.Connect(user, password, host, port, dbname)
		if err != nil {
			return
		}
	}

	handler.RegisDB(db)
	router := gin.New()
	routes.Router(router)
	router.Run(":8080")
}

func InitDatabase() (*sql.DB, error) {
	db, err := lib.Connect(user, password, host, port, "")
	if err != nil {

		return nil, err
	}
	// if err := lib.DropDB(db, dbname); err != nil {
	// 	return nil, err
	// }

	if err := lib.CreateDatabase(db, dbname); err != nil {
		return nil, err
	}

	db, err = lib.Connect(user, password, host, port, dbname)
	if err != nil {
		return nil, err
	}

	if err = lib.Use(db, dbname); err != nil {
		return nil, err
	}

	if err = lib.CreateTable(db, controler.TbStock); err != nil {
		return nil, err
	}

	data := []*controler.Stock{
		{Product: "Mie Goreng", Location: "A-1-1", Quantity: 100, CreateDate: time.Now()},
		{Product: "Kopi", Location: "A-1-2", Quantity: 100, CreateDate: time.Now()},
	}

	for _, item := range data {
		if err := item.Insert(db); err != nil {
			return nil, err
		}
	}

	if err = lib.CreateTable(db, controler.Tblog); err != nil {
		return nil, err
	}

	return db, err
}
