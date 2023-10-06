package model

import "time"

type Audit struct {
	ID        uint `gorm:"primaryKey"`
	Userid    string
	Operation string
	Metadata  string
	Timestamp time.Time
}
