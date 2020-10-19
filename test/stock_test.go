package test

import (
	"database/sql"
	"fmt"
	"klik/controler"
	"testing"
	"time"

	"github.com/mnurfilza/lib"
)

var user, password, host, port, dbname string

func init() {
	user = "root"
	password = "mypass123" // adjust with your mysql configuration
	host = "172.17.0.2"    //adjust with your configuration
	port = "3306"
	dbname = "klik_test"
}
func InitDb() (*sql.DB, error) {
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

	if err = lib.CreateTable(db, controler.Tblog); err != nil {
		return nil, err
	}

	return db, err
}

func TestStock(t *testing.T) {
	data := []*controler.Stock{
		{Product: "Mie Goreng", Location: "A-1-1", Quantity: 100, CreateDate: time.Now()},
		{Product: "Kopi", Location: "A-1-2", Quantity: 100, CreateDate: time.Now()},
	}

	db, err := InitDb()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	t.Run("Test Insert and Get Untuk Tb Berat", func(t *testing.T) {
		for _, item := range data {
			if err := item.Insert(db); err != nil {
				t.Fatal(err)
			}
		}

		fmt.Println("Insert Sukses")
	})

	t.Run("Testing Update Table Berat", func(t *testing.T) {
		update := map[string]interface{}{
			"adjustment": -15,
		}
		data := controler.Stock{ID: 2}
		_, err := data.Update(db, update)
		if err != nil {
			t.Fatal(err)
		}

		if err := data.Get(db); err != nil {
			t.Fatal(err)
		}

	})
}
