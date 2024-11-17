package utils

import (
	"time"
)

var JST *time.Location

func InitTimezone() error {
	var err error
	JST, err = time.LoadLocation("Asia/Tokyo")
	return err
}