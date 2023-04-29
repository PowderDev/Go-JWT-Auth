package config

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"auth/dataservice/migrations"
	"gorm.io/gorm"
)

var envVars = []string{
	"JWT_REFRESH_SECRET",
	"JWT_ACCESS_SECRET",
	"DSN",
	"DOMAIN",
}

func CheckForMigration(db *gorm.DB) {
	migration := flag.Bool("migrate", false, "whether or not to run db migration")
	flag.Parse()

	if *migration {
		migrations.Migrate(db)
		return
	}
}

func CheckENV() error {
	for _, envVar := range envVars {
		envStr := os.Getenv(envVar)

		if envStr == "" {
			return errors.New(fmt.Sprintf("'%s'environment variable was not prodived", envVar))
		}
	}

	return nil
}
