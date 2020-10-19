package controler

import (
	"database/sql"
	"time"

	"github.com/mnurfilza/lib"
)

type Log struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	LocationID int       `json:"location_id"`
	Adjustment string    `json:"adjustment"`
	Quantity   int       `json:"quantity"`
	CreateDate time.Time `json:"createDate"`
}

var Tblog = `CREATE TABLE log (
	id int PRIMARY KEY AUTO_INCREMENT,
	type VARCHAR(55),
	locationID int,
	adjustment VARCHAR(55),
	quantity NUMERIC,
	createDate DATE);
	`

func (l *Log) Name() string {
	return "log"
}

func (l *Log) Fields() (fields []string, dst []interface{}) {
	fields = []string{"id", "type", "location_id", "adjustment", "quantity", "createDate"}
	dst = []interface{}{&l.ID, &l.Type, &l.LocationID, &l.Adjustment, &l.Quantity, &l.CreateDate}
	return fields, dst
}

func (l *Log) PrimaryKey() (fields []string, dst []interface{}) {
	fields = []string{"id"}
	dst = []interface{}{&l.ID}
	return fields, dst
}

func (l *Log) Structur() lib.Table {
	return &Log{}
}

func (l *Log) Insert(db *sql.DB) error {
	l.CreateDate = time.Now()
	return lib.Insert(db, l)
}

func (l *Log) Get(db *sql.DB) error {
	return lib.Get(db, l)
}

func (l *Log) Update(db *sql.DB, change map[string]interface{}) (map[string]interface{}, error) {
	return change, lib.Update(db, l, change)
}

func (l *Log) Delete(db *sql.DB) error {
	return lib.Delete(db, l)
}

func Logs(db *sql.DB, params lib.RequestParams) ([]*Log, error) {
	l := &Log{}
	res, err := lib.Fetch(db, l, params)
	if err != nil {
		return nil, err
	}
	mhs := make([]*Log, len(res))
	for index, item := range res {
		mhs[index] = item.(*Log)
	}
	return mhs, nil
}
