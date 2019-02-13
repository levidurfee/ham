package log

// Entry of contact between two hams
type Entry struct {
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
}
