// Package database provide method to manage in-memory database
package database

import (
	"graylog-alert-exporter/internal/utils"

	"github.com/hashicorp/go-memdb"
	"github.com/sirupsen/logrus"
)

const TableName = "alert"
const IndexName = "id"

var db *memdb.MemDB

// Alert is struct of database schema
type Alert struct {
	ID      string            `json:"ID"`
	Timeout int               `json:"Timeout"`
	Data    map[string]string `json:"Data"`
}

// Init is function to create database schema and create persistance database connection
func Init() {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			TableName: {
				Name: TableName,
				Indexes: map[string]*memdb.IndexSchema{
					IndexName: {
						Name:         IndexName,
						AllowMissing: false,
						Unique:       true,
						Indexer:      &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}
	var err error
	if db, err = memdb.NewMemDB(schema); err != nil {
		logrus.Fatal(err)
	}
}

// InsertAlert will add Alert into database
func InsertAlert(alert Alert) error {
	logrus.Debugf("Insert alert to database\n%s\n", utils.PrettyJSON(alert))
	txn := db.Txn(true)
	if err := txn.Insert(TableName, alert); err != nil {
		logrus.Error("Error to insert record into database: ", err)
		return err
	}
	txn.Commit()
	return nil
}

// InsertAlerts will add slice of Alert into database
func InsertAlerts(alerts []Alert) {
	txn := db.Txn(true)
	for i, alert := range alerts {
		logrus.Debugf("Insert alert %d to database\n%s\n", i, utils.PrettyJSON(alert))
		if err := txn.Insert(TableName, alert); err != nil {
			logrus.Error("Error to insert records into database: ", err)
			return
		}
	}
	txn.Commit()
}

// GetAlertByID return Alert that match with Title from database
func GetAlertByID(title string) Alert {
	txn := db.Txn(false)
	raw, err := txn.First(TableName, IndexName, title)
	if err != nil {
		logrus.Error("Error to get record by title from database: ", err)
		return Alert{}
	}
	logrus.Debugf("Get alert by name %s\n%v\n", title, utils.PrettyJSON(raw.(Alert)))
	return raw.(Alert)
}

// GetAllAlerts return all alerts in database as slice of struct
func GetAllAlerts() []Alert {
	txn := db.Txn(false)
	it, err := txn.Get(TableName, IndexName)
	if err != nil {
		logrus.Error("Error to get all records from database: ", err)
		return nil
	}
	alerts := make([]Alert, 0)
	for obj := it.Next(); obj != nil; obj = it.Next() {
		alerts = append(alerts, obj.(Alert))
	}
	logrus.Debugf("Get all alerts\n%v\n", utils.PrettyJSON(alerts))
	return alerts
}

// RemoveAlert remove a record by id
func RemoveAlert(id string) error {
	txn := db.Txn(true)
	if err := txn.Delete(TableName, GetAlertByID(id)); err != nil {
		return err
	}
	txn.Commit()
	return nil
}
