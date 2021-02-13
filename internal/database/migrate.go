package database

import (
	"context"
	"log"
	"os"
	"time"

	authorizedModel "github.com/duyquang6/go-rbac-practice/internal/authorization/model"
	userModel "github.com/duyquang6/go-rbac-practice/internal/user/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Close releases database connections.
func (_db *DB) Migrate(ctx context.Context) error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	db := _db.Pool.Session(&gorm.Session{Logger: newLogger})

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "00001-CreateRole",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&authorizedModel.Role{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("roles")
			},
		},
		{
			ID: "00002-CreatePermission",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&authorizedModel.Permission{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("permissions")
			},
		},
		{
			ID: "00003-CreatePolicy",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&authorizedModel.Policy{})
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("policy_roles"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("permission_policies"); err != nil {
					return err
				}
				return tx.Migrator().DropTable("policies")
			},
		},
		{
			ID: "00004-CreateUser",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&userModel.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("user_roles"); err != nil {
					return err
				}
				return tx.Migrator().DropTable("users")
			},
		},
		{
			ID: "00005-AddUserPermission",
			Migrate: func(tx *gorm.DB) error {
				statements := []string{
					`INSERT INTO permissions(object,code,action) VALUES ('user', 1, 'read'),
																	('user', 2, 'create'),
																	('user', 4, 'update'),
																	('user', 8, 'delete');`,
				}
				for _, sql := range statements {
					if err := tx.Exec(sql).Error; err != nil {
						return err
					}
				}
				return nil
			},
		},
	})
	return m.Migrate()
}
