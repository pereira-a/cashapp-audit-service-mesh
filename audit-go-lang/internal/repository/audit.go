package repository

import (
	"time"

	"gorm.io/gorm"
)

type Audit struct {
	ID        uint `gorm:"primaryKey"`
	Userid    string
	Operation string
	Metadata  string
	Timestamp time.Time
}

type AuditRepository struct {
	DB_CONN *gorm.DB
}

func New() *AuditRepository {
	println(DB)
	return &AuditRepository{
		DB_CONN: DB,
	}
}

func (r AuditRepository) FindByUserId(userid string) []Audit {
	// println(test)
	var audits []Audit
	r.DB_CONN.Where("userid = ?", userid).Find(&audits)
	return audits
}

func (r AuditRepository) Create() string {
	// TODO
	return ""
}
