package utils

import (
	"strconv"
	"time"
)

func ConvertUint(s string) uint64 {
	parseUint, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return parseUint
}

func ConvertFloat(s string) float64 {
	parseFloat, err := strconv.ParseFloat(s, 10)
	if err != nil {
		return 0
	}
	return parseFloat
}

func ConvertToString(n uint8) string {
	return strconv.FormatUint(uint64(n), 10)
}

func GetNowString() string {
	t := time.Now()
	return t.Format("2006-01-02")
}
