package models

type Zone struct {
	Name string `json:"name" bson:"name"`

	// ID is the CloudStack zone ID
	// Deprecated: Use Name instead
	ID string `json:"-" bson:"-"`
}
