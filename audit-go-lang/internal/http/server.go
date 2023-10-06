package http

import (
	"audit/internal/model"
	"audit/internal/repository"
	"audit/internal/service"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Audit struct {
	Userid    string
	Operation string
	Metadata  string `json:",omitempty"`
	Timestamp time.Time
}

type auditService interface {
	GetAuditsByUserId(userid string) []model.Audit
	CreateAudit(model model.Audit) string
}

var logic auditService = service.New(repository.New())

func ListenAndServe() {
	e := echo.New()

	e.GET("/audit", GetAudit)
	e.POST("/audit", CreateAudit)

	e.Logger.Fatal(e.Start(":8080"))
}

func GetAudit(c echo.Context) error {
	audit := Audit{
		Userid:    "testUser",
		Operation: "sendMoney",
		Timestamp: time.Now(),
	}

	logic.GetAuditsByUserId("test")

	userId := c.QueryParam("userid")
	fmt.Printf("Requested audit for userId=%s\n", userId)
	return c.JSON(http.StatusOK, audit)
}

func CreateAudit(c echo.Context) error {
	audit := new(Audit)
	if err := c.Bind(audit); err != nil {
		return err
	}

	fmt.Printf("Received Audit: %s - %s - %s - %s\n", audit.Userid, audit.Operation, audit.Metadata, audit.Timestamp)
	return c.JSON(http.StatusCreated, audit)
}
