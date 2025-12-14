// service/pingSerivce.go
package service

import (
	"account-guru/internal/repository"
)

func PingService() error {
	// service is alive if this function is running
	if err := repository.PingDatabase(); err != nil {
		return err
	}
	return nil
}
