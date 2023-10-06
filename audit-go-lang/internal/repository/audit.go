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
	DB *gorm.DB
}

func New() *AuditRepository {
	return &AuditRepository{
		DB: DB,
	}
}

func (r AuditRepository) FindByUserId(userid string) []Audit {
	var audits []Audit
	r.DB.Where("userid = ?", userid).Find(&audits)
	return audits
}

func (r AuditRepository) Create() string {
	// TODO
	return ""
}
