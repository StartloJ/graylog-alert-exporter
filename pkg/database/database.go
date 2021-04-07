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
	ID          string
	Title       string
	Description string
	Source      string
	Priority    int
	Timeout     int
}

// Init is function to create database schema and create presistance database connection
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
func InsertAlert(alert Alert) {
	logrus.Debugf("Insert alert to database\n%s\n", utils.PrettyJSON(alert))
	txn := db.Txn(true)
	if err := txn.Insert(TableName, alert); err != nil {
		logrus.Fatal(err)
	}
	txn.Commit()
}

// InsertAlerts will add slice of Alert into database
func InsertAlerts(alerts []Alert) {
	txn := db.Txn(true)
	for i, alert := range alerts {
		logrus.Debugf("Insert alert %d to database\n%s\n", i, utils.PrettyJSON(alert))
		if err := txn.Insert(TableName, alert); err != nil {
			logrus.Fatal(err)
		}
	}
	txn.Commit()
}

// GetAlertByTitle return Alert thaht match with Title from database
func GetAlertByTitle(title string) Alert {
	txn := db.Txn(false)
	raw, err := txn.First(TableName, IndexName, title)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Debugf("Get alert by name %s\n%v\n", title, utils.PrettyJSON(raw.(Alert)))
	return raw.(Alert)
}

// GetAllAlerts return all alerts in database as slice of struct
func GetAllAlerts() (alerts []Alert) {
	txn := db.Txn(false)
	it, err := txn.Get(TableName, IndexName)
	if err != nil {
		logrus.Fatal(err)
	}
	for obj := it.Next(); obj != nil; obj = it.Next() {
		alerts = append(alerts, obj.(Alert))
	}
	logrus.Debugf("Get all alerts\n%v\n", utils.PrettyJSON(alerts))
	return alerts
}
