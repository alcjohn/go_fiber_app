package utils

import (
	"database/sql/driver"
	"time"
)

// AnyTime ..
type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
// https://github.com/DATA-DOG/go-sqlmock/blob/master/README.md#matching-arguments-like-timetime
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
