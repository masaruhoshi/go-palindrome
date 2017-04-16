package main

import (
	"testing"
)

func TestNewDaoToReturnObject(t *testing.T) {
	settings := GetSettings()

	dao := NewDao(settings)
	ExpectNotNil(t, dao)	
}

func TestNewDaoToReturnNil(t *testing.T) {
	settings := GetSettings()
	settings.HostName = "invalid.host"

	defer func() {
		r := recover()
		ExpectNotNil(t, r)
	}()
	
	NewDao(settings)
}