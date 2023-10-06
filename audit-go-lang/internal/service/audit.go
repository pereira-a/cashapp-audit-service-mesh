package service

import (
	"audit/internal/model"
	"audit/internal/repository"
	"fmt"
)

type AuditService struct {
	repo auditRepository
}

type auditRepository interface {
	Create() string
	FindByUserId(userid string) []repository.Audit
}

func New(auditRepo auditRepository) *AuditService {
	s := AuditService{
		repo: auditRepo,
	}
	return &s
}

func (s AuditService) GetAuditsByUserId(userid string) []model.Audit {
	// s.repo.FindByUserId(userid)
	fmt.Println("Test this organization!")
	return nil
}

func (s AuditService) CreateAudit(model model.Audit) string {
	// TODO
	return ""
}
