// repository/pingRepository.go
package repository

import (
	"context"
	"time"

	"account-guru/database"
)

func PingDatabase() error {
	sqlDB, err := database.DB.DB()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return sqlDB.PingContext(ctx)
}

