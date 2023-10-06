package model

import "time"

type Audit struct {
	Userid    string
	Operation string
	Metadata  string
	Timestamp time.Time
}
