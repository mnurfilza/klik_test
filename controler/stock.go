package controler

import (
	"database/sql"
	"math"
	"strconv"
	"time"

	"github.com/mnurfilza/lib"
)

type Stock struct {
	ID         int       `json:"id"`
	Product    string    `json:"product"`
	Location   string    `json:"location"`
	Quantity   int       `json:"quantity"`
	Adjustment int       `json:"adjustment"`
	CreateDate time.Time `json:"createDate"`
}

var TbStock = `CREATE TABLE stock (
	id int PRIMARY KEY AUTO_INCREMENT,
	product VARCHAR(55),
	location VARCHAR(55),
	quantity NUMERIC,
	createDate DATE);
	`

func (s *Stock) Name() string {
	return "stock"
}

func (s *Stock) Fields() (fields []string, dst []interface{}) {
	fields = []string{"id", "product", "location", "quantity", "createDate"}
	dst = []interface{}{&s.ID, &s.Product, &s.Location, &s.Quantity, &s.CreateDate}
	return fields, dst
}

func (s *Stock) PrimaryKey() (fields []string, dst []interface{}) {
	fields = []string{"id"}
	dst = []interface{}{&s.ID}
	return fields, dst
}

func (s *Stock) Structur() lib.Table {
	return &Stock{}
}

func (s *Stock) Insert(db *sql.DB) error {
	return lib.Insert(db, s)
}

func (s *Stock) Get(db *sql.DB) error {
	return lib.Get(db, s)
}

func (s *Stock) Update(db *sql.DB, change map[string]interface{}) (map[string]interface{}, error) {

	val, ok := change["adjustment"]
	if ok {
		log := &Log{}
		adj := math.Signbit(float64(val.(int)))
		log.Type = "Inbound"
		log.Adjustment = strconv.Itoa(val.(int))
		if adj {
			log.Adjustment = strconv.Itoa(val.(int))
			log.Type = "Outbound"
		}

		if err := s.Get(db); err != nil {
			return nil, err
		}

		log.Quantity = s.Quantity + val.(int)
		log.LocationID = s.ID
		if err := log.Insert(db); err != nil {
			return nil, err
		}
		delete(change, "adjustment")
		change["quantity"] = log.Quantity
	}

	return change, lib.Update(db, s, change)
}

func (s *Stock) Delete(db *sql.DB) error {
	return lib.Delete(db, s)
}

func Stocks(db *sql.DB, params lib.RequestParams) ([]*Stock, error) {
	s := &Stock{}
	res, err := lib.Fetch(db, s, params)
	if err != nil {
		return nil, err
	}
	mhs := make([]*Stock, len(res))
	for index, item := range res {
		mhs[index] = item.(*Stock)
	}
	return mhs, nil
}
