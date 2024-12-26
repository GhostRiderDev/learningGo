package model 

// RecordID defines a record id. Together with RecordType
// Identifies unique record accross all types.
type RecordID string


// RecordType defines a record type.  Together with RecordID
// Identifies unique record accross all types.
type RecordType string

// Existing record types.
const (
	RecordTypeMovie = RecordType("movie")
)

// UserID defines a user id.
type UserID string

// RatingValue defines a value of a rating record.
type RatingValue uint8

type Rating struct {
	RecordID string `json:"recordId"`
	RecordType string `json:"recordType"`
	UserID UserID `json:"userId"`
	Value RatingValue `json:"value"`
}