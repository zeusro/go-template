package audit

import (
	"time"

	"gorm.io/gorm"
)

// AuditLog represents an audit log entry
type AuditLog struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID    *uint   `json:"user_id"`
	Action    string  `gorm:"not null;index" json:"action"` // CREATE, UPDATE, DELETE, READ, etc.
	Resource  string  `gorm:"not null;index" json:"resource"` // user, order, etc.
	ResourceID *uint   `json:"resource_id"`
	IPAddress string  `json:"ip_address"`
	UserAgent string  `json:"user_agent"`
	Details   string  `gorm:"type:text" json:"details"` // JSON string with additional details
	Status    string  `gorm:"default:success" json:"status"` // success, failed
	Error     *string `json:"error,omitempty"`
}

// TableName specifies the table name
func (AuditLog) TableName() string {
	return "audit_logs"
}
