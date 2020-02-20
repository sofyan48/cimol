package v1

import "time"

// Logging ...
type Logging struct {
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	TimeAt      time.Time   `json:"time"`
}
