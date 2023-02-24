package DBRequests

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"time"
)

type Entry struct {
	Id        int64
	SessionId string
	Frequency float64
	Timestamp time.Time
}

func (u Entry) String() string {
	return fmt.Sprintf("Entry# %d\nSessionId: %s\nFrequency: %f\nTimestamp: %s\n",
		u.Id, u.SessionId, u.Frequency, u.Timestamp)
}

func CreateConnection(port string, user string, password string, database string) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     port,
		User:     user,
		Password: password,
		Database: database,
	})
	return db
}

func CreateTable(db *pg.DB) error {

	tableModel := (*Entry)(nil)

	err := db.Model(tableModel).CreateTable(&orm.CreateTableOptions{})
	if err != nil {
		return err
	}
	return nil
}

func InsertElem(db *pg.DB, elem *Entry) error {
	_, err := db.Model(elem).Insert()
	if err != nil {
		return err
	}
	return nil
}

func SelectElems(db *pg.DB) (error, []Entry) {
	var elems []Entry
	err := db.Model(&elems).Select()
	if err != nil {
		return err, nil
	}
	return nil, elems
}

func AddEntry(db *pg.DB, SessionId string, Frequency float64, Timestamp time.Time) error {
	elem := &Entry{
		SessionId: SessionId,
		Frequency: Frequency,
		Timestamp: Timestamp,
	}
	err := InsertElem(db, elem)
	if err != nil {
		return err
	}
	return nil
}
