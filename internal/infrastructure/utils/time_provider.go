package utils

import (
	"time"
)

type Provider interface {
	Now() time.Time
}

type RealTimeProvider struct{}

func NewRealTimeProvider() *RealTimeProvider {
	return &RealTimeProvider{}
}

func (RealTimeProvider) Now() time.Time {
	return time.Now()
}
