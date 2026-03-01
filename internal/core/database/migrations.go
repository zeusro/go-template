package database

import (
	"github.com/zeusro/go-template/internal/core/logprovider"
	"github.com/zeusro/go-template/internal/domain/entity"
	"github.com/zeusro/go-template/internal/domain/entity/audit"
)

// AutoMigrate runs database migrations
func AutoMigrate(db *DB, log logprovider.Logger) error {
	// Migrate domain entities
	if err := db.AutoMigrate(
		&entity.User{},
		&audit.AuditLog{},
		// Add more entities here
	); err != nil {
		log.Errorf("Database migration failed: %v", err)
		return err
	}

	log.Info("Database migrations completed successfully")
	return nil
}
