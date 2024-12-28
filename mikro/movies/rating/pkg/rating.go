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

// RatingEventType defines a type of a Rating record.
type RatingEventType string

// Rating event types.
const (
	RatingEventTypePut    = "put"
	RatingEventTypeDelete = "delete"
)

type Rating struct {
	RecordID   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}

type RatingEvent struct {
	RecordID   string          `json:"recordId"`
	RecordType string          `json:"recordType"`
	UserID     UserID          `json:"userId"`
	Value      RatingValue     `json:"value"`
	EventType  RatingEventType `json:"eventType"`
}
