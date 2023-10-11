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
	var eAudits = s.repo.FindByUserId(userid)
	var mAudits []model.Audit
	for _, eAudit := range eAudits {
		mAudits = append(mAudits, fromEntityToModel(eAudit))
	}
	fmt.Printf("Test this organization => %s", mAudits)
	return mAudits
}

func (s AuditService) CreateAudit(model model.Audit) string {
	// TODO
	return ""
}

func fromEntityToModel(entity repository.Audit) model.Audit {
	return model.Audit{
		Userid:    entity.Userid,
		Operation: entity.Operation,
		Metadata:  entity.Metadata,
		Timestamp: entity.Timestamp,
	}
}
