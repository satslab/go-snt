package search

import (
	"time"
)

type TimePeriod struct {
	From time.Time
	To   time.Time
}

type Query struct {
	Ingestion     TimePeriod
	BeginPosition TimePeriod
	EndPosition   TimePeriod
}
