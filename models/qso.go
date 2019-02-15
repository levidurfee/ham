package models

// QSOEntityType is for saving to the datastore
const QSOEntityType = "QSOEntry"

// QSO of contact between two hams
type QSO struct {
	UserID      string
	Date        string
	CallSign    string
	RSTSent     int
	RSTReceived int
	Frequency   float64
	Mode        string
	Power       string
	QTH         string
	Country     string
	Comments    string
	Band        int
	RequestID   int64
}
