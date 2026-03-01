package security

import (
	"context"
	"encoding/json"
	"time"

	"github.com/zeusro/go-template/internal/core/database"
	"github.com/zeusro/go-template/internal/core/logprovider"
	"github.com/zeusro/go-template/internal/domain/entity/audit"
)

// AuditLogger handles audit logging
type AuditLogger struct {
	db     *database.DB
	logger logprovider.Logger
}

// NewAuditLogger creates a new audit logger
func NewAuditLogger(db *database.DB, log logprovider.Logger) *AuditLogger {
	return &AuditLogger{
		db:     db,
		logger: log,
	}
}

// Log logs an audit event
func (a *AuditLogger) Log(ctx context.Context, entry *audit.AuditLog) error {
	if err := a.db.WithContext(ctx).Create(entry).Error; err != nil {
		a.logger.Errorf("Failed to log audit event: %v", err)
		return err
	}
	return nil
}

// LogAction logs a user action
func (a *AuditLogger) LogAction(ctx context.Context, userID *uint, action, resource string, resourceID *uint, ipAddress, userAgent string, details interface{}, status string, err error) error {
	var detailsJSON string
	if details != nil {
		data, e := json.Marshal(details)
		if e != nil {
			a.logger.Warnf("Failed to marshal audit details: %v", e)
		} else {
			detailsJSON = string(data)
		}
	}

	var errMsg *string
	if err != nil {
		msg := err.Error()
		errMsg = &msg
	}

	entry := &audit.AuditLog{
		UserID:     userID,
		Action:     action,
		Resource:   resource,
		ResourceID: resourceID,
		IPAddress:  ipAddress,
		UserAgent:  userAgent,
		Details:    detailsJSON,
		Status:     status,
		Error:      errMsg,
		CreatedAt:  time.Now(),
	}

	return a.Log(ctx, entry)
}
