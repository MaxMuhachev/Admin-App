package utils

import (
	"content/src/config"
	"fmt"
)

func FormatConnect(dns config.Dns) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", config.GetUserName(dns), config.GetPassword(dns), config.GetHostName(dns), config.GetDbName(dns))
}
