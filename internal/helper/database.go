package helper

import (
	"fmt"
	"golang-rest-api-articles/internal/config"
)

func GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.GetDbUser(),
		config.GetDbPassword(),
		config.GetDbHost(),
		config.GetDbPort(),
		config.GetDbName(),
	)
}
