package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func GetMigrations(db *gorm.DB) []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "Token_RefreshToken_Col_28_04",
			Migrate: func(tx *gorm.DB) error {
				type Token struct {
					RefreshToken string `json:"refreshToken"`
				}

				return tx.AutoMigrate(&Token{})
			},
			Rollback: func(tx *gorm.DB) error {
				return db.Migrator().DropTable("user_companies")
			},
		},
		{
			ID: "Initial_commit_27_04",
			Migrate: func(tx *gorm.DB) error {
				type Token struct {
					UserID       int    `json:"user_id" gorm:"primaryKey"'`
					refreshToken string `json:"refreshToken"`
				}

				type User struct {
					ID        int       `json:"id" gorm:"primaryKey"`
					Email     string    `json:"email"`
					Name      string    `json:"name"`
					Password  string    `json:"-"`
					Active    int       `json:"active"`
					CreatedAt time.Time `json:"created_at"`
					UpdatedAt time.Time `json:"updated_at"`
				}

				return db.AutoMigrate(&User{}, &Token{})
			},
			Rollback: func(tx *gorm.DB) error {
				return db.Migrator().DropTable("auth", "tokens", "companies")
			},
		},
	}
}
