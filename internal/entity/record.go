package entity

// struct record holds the information about a temperature record
type Record struct {
	// timestamp in RFC3339 format
	Timestamp string
	// temperature value
	Value float32
}
